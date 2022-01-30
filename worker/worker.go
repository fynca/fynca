package worker

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"time"

	"git.underland.io/ehazlett/finca"
	api "git.underland.io/ehazlett/finca/api/services/render/v1"
	"git.underland.io/ehazlett/finca/datastore"
	"git.underland.io/ehazlett/finca/pkg/profiler"
	"git.underland.io/ehazlett/finca/version"
	"github.com/dustin/go-humanize"
	"github.com/gogo/protobuf/jsonpb"
	minio "github.com/minio/minio-go/v7"
	miniocreds "github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/sirupsen/logrus"
)

var (
	// ERRGPUNotSupported is returned when the worker does not support GPU workloads
	ErrGPUNotSupported = errors.New("worker does not support GPU workloads")
)

type GPU struct {
	Vendor  string
	Product string
}

type Worker struct {
	id         string
	config     *finca.Config
	natsClient *nats.Conn
	ds         *datastore.Datastore
	stopCh     chan bool
	gpus       []*GPU
	gpuEnabled bool
	maxJobs    int
	jobLock    *sync.Mutex
}

func NewWorker(id string, cfg *finca.Config) (*Worker, error) {
	if v := cfg.ProfilerAddress; v != "" {
		p := profiler.NewProfiler(v)
		go func() {
			if err := p.Run(); err != nil {
				logrus.WithError(err).Errorf("error starting profiler on %s", v)
			}
		}()
	}
	nc, err := nats.Connect(cfg.NATSURL)
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to nats")
	}

	gpuEnabled, err := gpuEnabled()
	if err != nil {
		return nil, err
	}

	gpus, err := getGPUs()
	if err != nil {
		return nil, err
	}

	ds, err := datastore.NewDatastore(cfg)
	if err != nil {
		return nil, err
	}

	workerConfig := cfg.GetWorkerConfig(id)

	w := &Worker{
		id:         id,
		config:     cfg,
		natsClient: nc,
		ds:         ds,
		stopCh:     make(chan bool, 1),
		gpus:       gpus,
		gpuEnabled: gpuEnabled,
		maxJobs:    workerConfig.MaxJobs,
		jobLock:    &sync.Mutex{},
	}

	return w, nil
}

func (w *Worker) Run() error {
	if err := w.showWorkerInfo(); err != nil {
		return err
	}
	logrus.Infof("GPU enabled: %v", w.gpuEnabled)

	logrus.Debugf("connecting to nats: %s", w.config.NATSURL)

	js, err := w.natsClient.JetStream()
	if err != nil {
		return errors.Wrap(err, "error getting jetstream context")
	}

	// start worker heartbeat
	go w.workerHeartbeat()

	// start background listener for worker control
	go w.workerControlListener()

	msgCh := make(chan *nats.Msg, 1)
	doneCh := make(chan bool)

	// connect to nats stream and listen for messages
	logrus.Debugf("monitoring jobs on subject %s", w.config.NATSJobSubject)
	sub, err := js.PullSubscribe(w.config.NATSJobSubject, finca.WorkerQueueGroupName, nats.AckWait(w.config.GetJobTimeout()))
	if err != nil {
		return errors.Wrapf(err, "error subscribing to nats subject %s", w.config.NATSJobSubject)
	}

	logrus.Debugf("worker max jobs: %d", w.maxJobs)
	jobTickerInterval := time.Second * 5
	jobTicker := time.NewTicker(jobTickerInterval)

	jobsProcessed := 0
	go func() {
		defer func() {
			doneCh <- true
		}()

		for {
			select {
			case <-w.stopCh:
				logrus.Debug("unsubscribing from subject")
				sub.Unsubscribe()

				logrus.Debug("draining")
				sub.Drain()

				close(msgCh)
				return
			case <-jobTicker.C:
				msgs, err := sub.Fetch(1)
				if err != nil {
					// if stop has been called the subscription will be drained and closed
					// ignore the subscription error and exit
					if !sub.IsValid() {
						continue
					}
					if err == nats.ErrTimeout {
						// ignore NextMsg timeouts
						continue
					}
					logrus.Warn(err)
					continue
				}
				m := msgs[0]

				buf := bytes.NewBuffer(m.Data)
				workerJob := &api.WorkerJob{}
				if err := jsonpb.Unmarshal(buf, workerJob); err != nil {
					logrus.WithError(err).Error("error unmarshaling api.WorkerJob from message")
					continue
				}

				// report message is in progress
				m.InProgress(nats.AckWait(w.config.GetJobTimeout()))

				logrus.Debugf("processing job with timeout %s", w.config.GetJobTimeout())

				workerInfo, err := w.getWorkerInfo()
				if err != nil {
					logrus.WithError(err).Error("error getting worker info")
					continue
				}
				logrus.Debugf("worker info: %+v", workerInfo)

				ctx, cancel := context.WithTimeout(context.Background(), w.config.GetJobTimeout())

				result := &api.JobResult{}
				jobID := ""

				var pErr error
				switch v := workerJob.GetJob().(type) {
				case *api.WorkerJob_FrameJob:
					jobID = v.FrameJob.ID
					result, pErr = w.processFrameJob(ctx, v.FrameJob)
				case *api.WorkerJob_SliceJob:
					jobID = v.SliceJob.ID
					result, pErr = w.processSliceJob(ctx, v.SliceJob)
				default:
					logrus.Errorf("unknown job type %+T", v)
					cancel()
					continue
				}

				if pErr != nil {
					if pErr == ErrGPUNotSupported {
						m.Nak()
						cancel()
						continue
					}

					// publish error
					logrus.WithError(pErr).Errorf("error processing job")
					if cErr := ctx.Err(); cErr != nil {
						logrus.Warn("job timeout occurred during processing")
					} else {
						// no timeout; requeue job
						logrus.WithError(pErr).Error("error occurred while processing job")
						if result == nil {
							result = &api.JobResult{}
						}
						result.Status = api.JobStatus_ERROR
						result.Error = pErr.Error()
						cancel()
					}
				}

				logrus.Infof("completed job %s: status=%s", jobID, result.Status)

				// publish status event for server
				b := &bytes.Buffer{}
				if err := w.ds.Marshaler().Marshal(b, result); err != nil {
					logrus.WithError(err).Error("error publishing job result")
					continue
				}
				js.Publish(w.config.NATSJobStatusSubject, b.Bytes())

				// ack message for completion
				m.Ack()

				// check for max processed jobs
				jobsProcessed += 1
				if w.maxJobs != 0 && jobsProcessed >= w.maxJobs {
					logrus.Infof("worker reached max jobs (%d), exiting", w.maxJobs)
					sub.Unsubscribe()
					sub.Drain()
					close(msgCh)
					return
				}
			}
		}
	}()

	<-doneCh
	logrus.Debug("worker finished")

	return nil
}

func (w *Worker) Stop() error {
	w.jobLock.Lock()
	defer w.jobLock.Unlock()

	w.stopCh <- true
	return nil
}

func (w *Worker) getMinioClient() (*minio.Client, error) {
	return minio.New(w.config.S3Endpoint, &minio.Options{
		Creds:  miniocreds.NewStaticV4(w.config.S3AccessID, w.config.S3AccessKey, ""),
		Secure: w.config.S3UseSSL,
	})
}

func (w *Worker) getWorkerInfo() (*api.Worker, error) {
	gpuInfo := []string{}
	for _, gpu := range w.gpus {
		gpuInfo = append(gpuInfo, fmt.Sprintf("%s: %s", gpu.Vendor, gpu.Product))
	}

	cpus, err := cpu.Counts(true)
	if err != nil {
		return nil, err
	}

	loadStats, err := load.Avg()
	if err != nil {
		return nil, err
	}

	m, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return &api.Worker{
		Name:            w.id,
		Version:         version.BuildVersion(),
		CPUs:            uint32(cpus),
		MemoryTotal:     int64(m.Total),
		MemoryAvailable: int64(m.Available),
		GPUs:            gpuInfo,
		Load1:           loadStats.Load1,
		Load5:           loadStats.Load5,
		Load15:          loadStats.Load15,
	}, nil
}

func (w *Worker) showWorkerInfo() error {
	info, err := w.getWorkerInfo()
	if err != nil {
		return err
	}

	logrus.Infof("CPU: %d", info.CPUs)
	logrus.Infof("Memory: %s", humanize.Bytes(uint64(info.MemoryTotal)))

	for _, gpu := range w.gpus {
		logrus.Infof("Graphics Card: %s", gpu)
	}

	return nil
}

func (w *Worker) workerHeartbeat() {
	t := time.NewTicker(finca.WorkerTTL)
	for range t.C {
		workerInfo, err := w.getWorkerInfo()
		if err != nil {
			logrus.WithError(err).Error("error getting worker info")
			continue
		}

		ctx := context.Background()
		logrus.Debugf("updating worker info: %+v", workerInfo)
		if err := w.ds.UpdateWorkerInfo(ctx, workerInfo); err != nil {
			logrus.WithError(err).Error("error updating worker info")
			continue
		}
	}
}
