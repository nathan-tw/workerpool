package main

type WorkerPool struct {
	PoolSize    int
	tasksSize   int
	tasksChan    chan Task
	resultsChan chan Task
	Results     func() []Task
}

func NewWorkerPool(tasks []Task, size int) *WorkerPool {
	tasksChan, resultsChan := make(chan Task, len(tasks)), make(chan Task, len(tasks))
	for _, task := range tasks {
		tasksChan <- task
	}
	close(tasksChan)
	pool := &WorkerPool{
		PoolSize: size,
		tasksSize: len(tasks),
		tasksChan: tasksChan,
		resultsChan: resultsChan,
	}
	pool.Results = pool.results
	return pool
}
func (pool *WorkerPool) Start() {  
    for i := 0; i < pool.PoolSize; i++ {  
        go pool.worker()  
    }  
}  

func (pool *WorkerPool) worker() {  
    for task := range pool.tasksChan {  
        task.Err = task.Do()  
        pool.resultsChan <- task  
    }  
}  

func (wp *WorkerPool) results() []Task {
	tasks := make([]Task, wp.tasksSize)
	for i := 0; i < wp.tasksSize; i++ {
		tasks[i] = <-wp.resultsChan
	}
	return tasks
}
