package render

import (
	"encoding/json"

	"git.underland.io/ehazlett/finca"
	api "git.underland.io/ehazlett/finca/api/services/render/v1"
	nats "github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func (s *service) jobStatusListener() {
	js, err := s.natsClient.JetStream()
	if err != nil {
		logrus.Fatal(err)
	}

	sub, err := js.PullSubscribe(s.config.NATSJobStatusSubject, finca.ServerQueueGroupName)
	if err != nil {
		logrus.Fatal(err)
	}

	s.serverQueueSubscription = sub

	msgCh := make(chan *nats.Msg)

	go func() {
		for {
			msgs, err := sub.Fetch(1)
			if err != nil {
				// skip invalid sub errors when exiting
				if !sub.IsValid() {
					return
				}
				if err == nats.ErrTimeout {
					// ignore NextMsg timeouts
					continue
				}
				logrus.WithError(err).Warn("error getting message")
				continue
			}

			msgCh <- msgs[0]
		}
	}()

	for {
		select {
		case <-s.stopCh:
			return
		case m := <-msgCh:
			var job *api.Job
			if err := json.Unmarshal(m.Data, &job); err != nil {
				logrus.WithError(err).Error("error unmarshaling api.Job from message")
				continue
			}
			logrus.Infof("job %s [%s] (frame: %d slice: %d) completed in %s on worker %s: success: %v",
				job.ID,
				job.Request.Name,
				job.RenderFrame,
				job.RenderSliceIndex,
				job.Duration,
				job.Worker.Name,
				job.Succeeded,
			)
			// TODO: store to datastore for UI?
			m.Ack()
		}
	}
}
