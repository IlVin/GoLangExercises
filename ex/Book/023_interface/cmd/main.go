package main

import (
	"company_interface/pkg/company"
	"company_interface/pkg/person"
	"company_interface/pkg/robot"
)

func main() {
	pers := person.Person{
		Name:     "Ilia",
		LastName: "Vinokurov",
		Age:      51,
	}

	robot := robot.Robot{
		Id:        "RoBot",
		Model:     "R2D2",
		WorkCount: 0,
	}

	comp := company.Company{}

	comp.Hire(&robot)
	comp.Hire(&pers) // мы передаём переменную типа Person в функцию, аргументом которой является переменная Worker!
	comp.Hire(pers)  // мы передаём переменную типа Person в функцию, аргументом которой является переменная Worker!
	comp.DoWork()
	comp.DoWork()

	return
}
