package worker_pool

type WP interface {
	Push(task ...func())
}

type WorkerPool struct {
	taskQueue chan func()
	Result    chan interface{}
}

func New(workerNum int) *WorkerPool {
	wp := &WorkerPool{
		taskQueue: make(chan func()),
		Result:    make(chan interface{}),
	}

	for i := 0; i < workerNum; i++ {
		go NewWorker(wp.taskQueue, wp.Result)
	}

	return wp
}

func (wp *WorkerPool) Push(task ...func()) {
	for _, t := range task {
		wp.taskQueue <- t
	}
}
