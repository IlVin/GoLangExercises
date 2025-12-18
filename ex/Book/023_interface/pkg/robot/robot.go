package robot

import (
	"company_interface/pkg/task"
	"fmt"
)

type Robot struct {
	Id        string
	Model     string
	WorkCount int
}

func (r Robot) String() string {
	return fmt.Sprintf("Robot %s model %s", r.Id, r.Model)
}

func (r *Robot) Work(tasks task.TaskReader) string {
	r.WorkCount++
	var plan = make([]task.Task, 3)
	_, err := tasks.Read(plan)
	if err != nil {
		panic(err)
	}
	for _, t := range plan {
		fmt.Println(t)
	}
	return fmt.Sprintf("%s worked %d times", r, r.WorkCount)
}
