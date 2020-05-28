package threadpool

import "fmt"

// Thread thread struct
type Thread struct {
	JobChan chan interface{} // channel to push job.

	CloseSignal chan struct{}
}

// NewThread create one thread to excute jobs.
func NewThread(pool chan Thread, closeSinal chan struct{}) {
	t := Thread{}
	t.JobChan = make(chan interface{})
	t.CloseSignal = closeSinal
	go func() {
		for {
			pool <- t
			select {
			case job := <-t.JobChan:
				fmt.Println(job)
				// TODO: excute
			case <-t.CloseSignal:
				return
			}
		}
	}()
}
