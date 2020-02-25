package pool

import (
	"fmt"
	"github.com/chariot9/worker-pool/job"
	"github.com/chariot9/worker-pool/result"
	"sync"
	"time"
)

type Pool struct {
	numOfJobs int
	jobs      chan job.Job
	results   chan result.Result
	done      chan bool
	completed bool
}

func NewPool(numOfJobs int) *Pool {
	p := &Pool{numOfJobs: numOfJobs}
	p.jobs = make(chan job.Job, numOfJobs)
	p.results = make(chan result.Result, numOfJobs)
	return p
}

func (p *Pool) Start(resources []interface{}, pro1 job.ProcessJob, pro2 result.ProcessResult) {
	start := time.Now()

	go p.allocate(resources)
	p.done = make(chan bool)
	go p.collect(pro2)
	go p.workerPool(pro1)
	<-p.done
	end := time.Now()
	diff := end.Sub(start)
	fmt.Printf("total time taken: [%f] seconds\n", diff.Seconds())
}

func (p *Pool) allocate(jobs []interface{}) {
	defer close(p.jobs)
	for i, v := range jobs {
		j := job.Job{Id: i, Resource: v}
		p.jobs <- j
	}
}

func (p *Pool) work(wg *sync.WaitGroup, proc job.ProcessJob) {
	defer wg.Done()
	for j := range p.jobs {
		output := result.Result{Job: j, Err: proc(j.Resource)}
		p.results <- output
	}
}

func (p *Pool) workerPool(proc job.ProcessJob) {
	defer close(p.results)
	var wg sync.WaitGroup
	for i := 0; i < p.numOfJobs; i++ {
		wg.Add(1)
		go p.work(&wg, proc)
	}
	wg.Wait()
}

func (p *Pool) collect(proc result.ProcessResult) {
	for r := range p.results {
		_ = proc(r)
	}

	p.done <- true
	p.completed = true
}

func (p *Pool) IsCompleted() bool {
	return p.completed
}
