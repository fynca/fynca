// Copyright 2022 Evan Hazlett
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package datastore

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"time"

	"github.com/fynca/fynca"
	api "github.com/fynca/fynca/api/services/jobs/v1"
	"github.com/go-redis/redis/v8"
	"github.com/gogo/protobuf/proto"
	minio "github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	jobArchiveTTL = time.Second * 24 * 60 * 60
)

// CreateJobArchive creates a zip archive of the job and stores a signed url from the storage service
func (d *Datastore) CreateJobArchive(ctx context.Context, jobID string) error {
	namespace := ctx.Value(fynca.CtxNamespaceKey).(string)

	job, err := d.GetJob(ctx, jobID)
	if err != nil {
		return err
	}
	// check for rendering and if slices return composite
	if job.Status == api.JobStatus_RENDERING {
		return fmt.Errorf("cannot create archive: job %s is still rendering", job.ID)
	}

	// create initial job archive to signal in progress
	jobArchiveKey := getJobArchiveKey(namespace, job.ID)
	jobArchive := &api.JobArchive{}

	data, err := proto.Marshal(jobArchive)
	if err != nil {
		return err
	}
	if err := d.redisClient.Set(ctx, jobArchiveKey, data, 0).Err(); err != nil {
		return errors.Wrapf(err, "error saving job archive for %s in database", job.ID)
	}

	// get files and create archive
	renderPath := path.Join(namespace, fynca.S3RenderPath, job.ID)
	objCh := d.storageClient.ListObjects(ctx, d.config.S3Bucket, minio.ListObjectsOptions{
		Prefix:    renderPath,
		Recursive: true,
	})

	frameKeys := []string{}
	for o := range objCh {
		if o.Err != nil {
			return o.Err
		}

		// TODO: support other file types
		if filepath.Ext(o.Key) != ".png" {
			continue
		}
		// ignore slice renders
		sliceMatch, err := regexp.MatchString(".*_slice-\\d+_\\d+\\.[png]", o.Key)
		if err != nil {
			return err
		}
		if sliceMatch {
			continue
		}
		frameMatch, err := regexp.MatchString(".*_\\d+.[png]", o.Key)
		if err != nil {
			return err
		}
		if frameMatch {
			frameKeys = append(frameKeys, o.Key)
		}
	}

	if len(frameKeys) == 0 {
		return fmt.Errorf("unable to find render images for job %s", job.ID)
	}

	tmpArchive, err := os.CreateTemp("", "-fynca-job-archive-")
	if err != nil {
		return err
	}
	defer os.Remove(tmpArchive.Name())

	zw := zip.NewWriter(tmpArchive)
	start := time.Now()

	// get project source
	pObj, err := d.storageClient.GetObject(ctx, d.config.S3Bucket, job.GetJobSource(), minio.GetObjectOptions{})
	if err != nil {
		return errors.Wrapf(err, "error getting job source %s", job.GetJobSource())
	}

	pf, err := zw.Create(filepath.Base(job.GetJobSource()))
	if err != nil {
		return err
	}

	if _, err := io.Copy(pf, pObj); err != nil {
		return err
	}

	// get rendered frames
	for _, k := range frameKeys {
		// get key data
		obj, err := d.storageClient.GetObject(ctx, d.config.S3Bucket, k, minio.GetObjectOptions{})
		if err != nil {
			return errors.Wrapf(err, "error getting render object from storage %s", k)
		}

		f, err := zw.Create(filepath.Base(k))
		if err != nil {
			return errors.Wrapf(err, "error creating file in job archive for %s", job.ID)
		}

		if _, err := io.Copy(f, obj); err != nil {
			return err
		}
	}
	if err := zw.Close(); err != nil {
		return err
	}
	logrus.Debugf("job %s archive time: %s", job.ID, time.Now().Sub(start))

	archive, err := os.Open(tmpArchive.Name())
	if err != nil {
		return err
	}

	fi, err := archive.Stat()
	if err != nil {
		return err
	}

	// copy archive to storage
	objectPath := path.Join(namespace, fynca.S3ProjectPath, job.ID, fmt.Sprintf("%s.zip", job.ID))
	if _, err := d.storageClient.PutObject(ctx, d.config.S3Bucket, objectPath, archive, fi.Size(), minio.PutObjectOptions{ContentType: "application/zip"}); err != nil {
		return err
	}

	// get presigned url for archive
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s.zip\"", job.ID))

	// valid for 24 hours
	presignedURL, err := d.storageClient.PresignedGetObject(ctx, d.config.S3Bucket, objectPath, jobArchiveTTL, reqParams)
	if err != nil {
		return err
	}

	jobArchive.ArchiveUrl = presignedURL.String()

	jData, err := proto.Marshal(jobArchive)
	if err != nil {
		return err
	}
	if err := d.redisClient.Set(ctx, jobArchiveKey, jData, jobArchiveTTL).Err(); err != nil {
		return errors.Wrapf(err, "error saving job archive for %s in database", job.ID)
	}

	return nil
}

func (d *Datastore) GetJobArchiveStatus(ctx context.Context, jobID string) (*api.JobArchive, error) {
	namespace := ctx.Value(fynca.CtxNamespaceKey).(string)
	job, err := d.GetJob(ctx, jobID)
	if err != nil {
		return nil, err
	}

	key := getJobArchiveKey(namespace, job.ID)
	data, err := d.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "error getting job archive %s from database", jobID)
	}

	jobArchive := api.JobArchive{}
	if err := proto.Unmarshal(data, &jobArchive); err != nil {
		return nil, err
	}

	return &jobArchive, nil
}

func getJobArchiveKey(namespace, jobID string) string {
	return path.Join(dbPrefix, namespace, jobID, "archive")
}
