package worker_pool

type WP interface {
	Push(task ...func()) error
}

type WorkerPool struct {
	taskQueue chan Task
	Result    chan Result
}

func New(workerNum int) *WorkerPool {
	wp := &WorkerPool{
		taskQueue: make(chan Task, workerNum),
		Result:    make(chan Result, workerNum),
	}

	for range workerNum {
		go NewWorker(wp.taskQueue, wp.Result)
	}

	return wp
}

func (wp *WorkerPool) Push(task ...Task) {
	for _, t := range task {
		wp.taskQueue <- t
	}
}

// Close closes the task queue channel
func (wp *WorkerPool) Close() {
	close(wp.taskQueue)
}
