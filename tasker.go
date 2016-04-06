package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	titan_go "github.com/iron-io/titan_go"
)

type Tasker struct {
	api *titan_go.JobsApi
	log log.FieldLogger
}

// Titan tasker.
func NewTasker(config *Config, log log.FieldLogger) *Tasker {
	api := titan_go.NewJobsApiWithBasePath(config.ApiUrl)
	return &Tasker{api, log}
}

func (t *Tasker) Job() *titan_go.Job {
	l := t.log.WithField("action", "DequeueJob")
	var job *titan_go.Job
	for {
		jobs, err := t.api.JobsGet(1)
		if err != nil {
			l.WithError(err).Errorln("dequeue job from api")
		} else if len(jobs.Jobs) > 0 {
			job = &jobs.Jobs[0]
			break
		}
		time.Sleep(1 * time.Second)
	}
	return job
}

func (t *Tasker) Update(job *titan_go.Job) error {
	l := t.log.WithFields(log.Fields{
		"action": "UpdateJob",
		"job_id": job.Id,
	})
	l.Debugln("Sending PATCH to update job", job)
	j, err := t.api.GroupsGroupNameJobsIdPatch(job.GroupName, job.Id, titan_go.JobWrapper{*job})
	if err != nil {
		l.WithError(err).Errorln("Update failed")
		return err
	}
	l.Infoln("Got back", j)
	return nil
}

// TODO: this should be on server side
func (t *Tasker) RetryTask(job *titan_go.Job) error {
	panic("Not implemented Retry")
}

func (t *Tasker) IsCancelled(job *titan_go.Job) bool {
	wrapper, err := t.api.GroupsGroupNameJobsIdGet(job.GroupName, job.Id)
	if err != nil {
		log.Errorln("JobIdGet from Cancel", "err", err)
		return false
	}

	// FIXME(nikhil) Current branch does not capture cancellation.
	return wrapper.Job.Status == "error"
}

func (t *Tasker) Succeeded(job *titan_go.Job, r string) error {
	j, err := t.api.GroupsGroupNameJobsIdPatch(job.GroupName, job.Id, titan_go.JobWrapper{*job})
	if err != nil {
		log.Errorln("Update failed", "job", job.Id, "err", err)
		return err
	}
	log.Infoln("Got back", j)
	// _, err := t.api.JobIdSuccessPost(job.Id, r)
	// if err != nil {
	// 	log.Errorln("JobIdSuccessPost", "jobId", job.Id, "err", err)
	// }
	return nil
}

func (t *Tasker) Failed(job *titan_go.Job, reason string, r string) error {
	j, err := t.api.GroupsGroupNameJobsIdPatch(job.GroupName, job.Id, titan_go.JobWrapper{*job})
	if err != nil {
		log.Errorln("Update failed", "job", job.Id, "err", err)
		return err
	}
	log.Infoln("Got back", j)
	// _, err := t.api.JobIdFailPost(job.Id, reason, "" /* details */, r)
	// if err != nil {
	// 	log.Errorln("JobIdFailPost", "jobId", job.Id, "err", err)
	// }
	return nil
}
