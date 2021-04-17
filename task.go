package main

type Task struct {
	Id  int
	Err error
	f   func() error
}

func (task *Task) Do() error {
	return task.f()
}
