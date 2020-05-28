package threadpool

// RunnableJob run方法没有返回值
type RunnableJob interface {
	Run()
}

// CallableJob call方法返回一个结果
type CallableJob interface {
	Call() interface{}
}

type callableTask struct {
	job    CallableJob
	future *Future
}

// Future state of CallableJob.
type Future struct {
	res  chan interface{}
	done bool
}

// Get return the res of CallableJob.
func (f *Future) Get() interface{} {
	return <-f.res
}

// IsDone return the finished status of CallableJob.
func (f *Future) IsDone() bool {
	return f.done
}
