package threadpool

import "fmt"

// Thread thread struct
type Thread struct {
	JobChan chan Job // channel to push job.
}

// NewThread create one thread to excute jobs.
func NewThread(pool chan Thread) {
	t := Thread{}
	t.JobChan = make(chan Job)
	go func() {
		for {
			pool <- t
			select {
			case job := <-t.JobChan:
				fmt.Println(job)
				// TODO: excute
			}
		}
	}()
}
