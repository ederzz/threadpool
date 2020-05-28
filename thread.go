package threadpool

// Thread thread struct
type Thread struct {
	JobChan chan interface{} // channel to push job.

	CloseSignal chan struct{}
}

// NewThread create one thread to execute jobs.
func NewThread(pool chan Thread, closeSinal chan struct{}) {
	t := Thread{}
	t.JobChan = make(chan interface{})
	t.CloseSignal = closeSinal
	go func() {
		for {
			pool <- t
			select {
			case job := <-t.JobChan:
				t.executeJob(job)
			case <-t.CloseSignal:
				return
			}
		}
	}()
}

func (t Thread) executeJob(j interface{}) {
	switch task := j.(type) {
	case RunnableJob:
		task.Run()
		break
	case callableTask:
		res := task.job.Call()
		task.future.res <- res
		task.future.done = true
		break
	}
}
