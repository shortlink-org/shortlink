package worker_pool

type Task func() (any, error)

type Result struct {
	Value any
	Error error
}

type Worker struct {
	taskQueue <-chan Task
	result    chan<- Result
}

func NewWorker(taskQueue <-chan Task, result chan<- Result) *Worker {
	w := &Worker{
		taskQueue: taskQueue,
		result:    result,
	}

	go w.run()

	return w
}

func (w *Worker) run() {
	for task := range w.taskQueue {
		result, err := task()
		w.result <- Result{result, err}
	}
}
