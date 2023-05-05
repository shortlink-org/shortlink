package worker_pool

type worker struct {
	taskQueue chan func()
	result    chan interface{}
}

func NewWorker(taskQueue chan func(), result chan interface{}) *worker {
	w := &worker{
		taskQueue: taskQueue,
		result:    result,
	}

	go w.run()

	return w
}

func (w *worker) run() {
	for {
		task := <-w.taskQueue
		task()
		w.result <- struct{}{}
	}
}
