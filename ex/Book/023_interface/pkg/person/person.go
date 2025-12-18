package person

import (
	"company_interface/pkg/task"
	"fmt"
)

type Person struct {
	Name     string
	LastName string
	Age      int
}

func (p Person) String() string {
	return fmt.Sprintf("Person: %s %s; Age %d", p.Name, p.LastName, p.Age)
}

func (p Person) Work(_ task.TaskReader) string {
	return fmt.Sprintf("Worker %s %s is working", p.Name, p.LastName)
}
