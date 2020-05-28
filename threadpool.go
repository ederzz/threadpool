package threadpool

import (
	"errors"
)

var errJobPoolFull = errors.New("Job's pool if full, can not add new Job.")

// ThreadPool thread-pool struct
type ThreadPool struct {
	JobSize    int
	ThreadSize int
	Jobs       chan Job    // pool of jobs.
	Threads    chan Thread // pool of threads.
}

// NewThreadPool create thread-pool
func NewThreadPool(poolSize int, jobSize int) *ThreadPool {
	tp := &ThreadPool{
		ThreadSize: poolSize,
		JobSize:    jobSize,
	}
	tp.Jobs = make(chan Job, jobSize)
	tp.Threads = make(chan Thread, poolSize)
	tp.initThreads()

	return tp
}

// init threads of thread-pool.
func (tp ThreadPool) initThreads() {
	for i := 0; i < tp.ThreadSize; i++ {
		NewThread(tp.Threads)
	}
}

// ExcuteJob excute a new job.
func (tp ThreadPool) ExcuteJob(job Job) error {
	if len(tp.Jobs) == tp.JobSize {
		return errJobPoolFull
	}
	tp.Jobs <- job
	return nil
}

// Start start to handle jobs.
func (tp ThreadPool) Start() {
	for {
		select {
		case job := <-tp.Jobs:
			thread := <-tp.Threads
			thread.JobChan <- job
			// TODO: case close
		}
	}
}

// Close close to handle jobs.
func (tp ThreadPool) Close() {
	// TODO: 关闭处理
	close(tp.JObs)
	close(tp.Threads)
}

// TODO: test && example
