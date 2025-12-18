package task

import "fmt"

type TaskReader interface {
	Read([]Task) (n int, err error)
}

type Task struct {
	Id int
}

type taskgen struct {
	NextTask int
}

func NewTaskReader() TaskReader {
	return &taskgen{NextTask: 3}
}

func (t *taskgen) Read(tasks []Task) (n int, err error) {
	for n < cap(tasks) {
		tasks[n] = Task{Id: t.NextTask}
		t.NextTask++
		n++
	}
	return n, nil
}

func (t *Task) String() string {
	return fmt.Sprintf("Task{%d}", t.Id)
}
