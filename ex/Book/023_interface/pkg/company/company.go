package company

import (
	"company_interface/pkg/task"
	"fmt"
)

type Worker interface {
	Work(tasks task.TaskReader) string
}

type Company struct {
	CompanyName string
	workers     []Worker
	tasks       task.TaskReader
}

func (c *Company) Hire(w Worker) string {
	c.workers = append(c.workers, w)
	return fmt.Sprintf("Company %s hire %s", c.CompanyName, w)
}

func (c *Company) DoWork() {
	if c.tasks == nil {
		c.tasks = task.NewTaskReader()
	}
	for _, w := range c.workers {
		fmt.Println(w.Work(c.tasks))
	}
}
