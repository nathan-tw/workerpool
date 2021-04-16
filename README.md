# workerpool

Simple worker pool practice in golang using buffered channel and goroutine. Worker pool is a task solving pattern.

<img src="https://img.cntofu.com/book/note/linux_system/images/580px-Thread_pool.svg.png" width="400px">


## Usage
```go
func main() {
    t := time.Now()  
  
    tasks := []Task{  
        {Id: 0, f: func() error { time.Sleep(2 * time.Second); fmt.Println(0); return nil }},  
        {Id: 1, f: func() error { time.Sleep(time.Second); fmt.Println(1); return errors.New("error") }},  
        {Id: 2, f: func() error { fmt.Println(2); return errors.New("error") }},  
    }  
    pool := NewWorkerPool(tasks, 2)  
    pool.Start()  
  
    tasks = pool.Results()  
    fmt.Printf("all tasks finished, timeElapsed: %f s\n", time.Now().Sub(t).Seconds())  
    for _, task := range tasks {  
        fmt.Printf("result of task %d is %v\n", task.Id, task.Err)  
    }  
}
```
