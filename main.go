// main.go

package main

import (
	"fmt"
	"time"
	"errors"
)


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
    fmt.Printf("all tasks finished, timeElapsed: %f s\n", time.Since(t).Seconds())  
    for _, task := range tasks {  
        fmt.Printf("result of task %d is %v\n", task.Id, task.Err)  
    }  
}
