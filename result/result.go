package result

import job "github.com/chariot9/worker-pool/job"

type Result struct {
	Job job.Job
	Err error
}

type ProcessResult func(result Result) error
