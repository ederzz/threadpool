package threadpool

import (
	"errors"
)

var errJobPoolFull = errors.New("Job's pool if full, can not add new Job.")

// ThreadPool thread-pool struct
type ThreadPool struct {
	JobSize    int
	ThreadSize int
	Jobs       chan interface{} // pool of jobs.
	Threads    chan Thread      // pool of threads.

	CloseSignal chan struct{}
}

// NewThreadPool create thread-pool
func NewThreadPool(poolSize int, jobSize int) *ThreadPool {
	tp := &ThreadPool{
		ThreadSize: poolSize,
		JobSize:    jobSize,
	}
	tp.Jobs = make(chan interface{}, jobSize)
	tp.Threads = make(chan Thread, poolSize)
	tp.initThreads()

	return tp
}

// init threads of thread-pool.
func (tp *ThreadPool) initThreads() {
	for i := 0; i < tp.ThreadSize; i++ {
		NewThread(tp.Threads, tp.CloseSignal)
	}
}

// push a job to threadpool.Jobs channel.
func (tp *ThreadPool) pushJob(job interface{}) error {
	if len(tp.Jobs) == tp.JobSize {
		return errJobPoolFull
	}
	tp.Jobs <- job
	return nil
}

// ExcuteRunnableJob excute a runnableJob.
func (tp *ThreadPool) ExcuteRunnableJob(job RunnableJob) error {
	return tp.pushJob(job)
}

// ExcuteCallableJob excute a callableJob.
func (tp *ThreadPool) ExcuteCallableJob(job CallableJob) (*Future, error) {
	future := &Future{}
	future.res = make(chan interface{})
	task := callableTask{job, future}
	err := tp.pushJob(task)
	if err != nil {
		return nil, err
	}

	return task.future, nil
}

// Start start to handle jobs.
func (tp *ThreadPool) Start() {
	for {
		select {
		case job := <-tp.Jobs:
			thread := <-tp.Threads
			thread.JobChan <- job
		case <-tp.CloseSignal:
			return
		}
	}
}

// Close close to handle jobs.
func (tp *ThreadPool) Close() {
	close(tp.CloseSignal)
	close(tp.Jobs)
	close(tp.Threads)
}

// TODO: test && example
// TODO: 添加指针
