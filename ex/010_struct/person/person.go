package person

import (
	"encoding/json"
	"time"
)

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func NewPerson(name string, email string) (Person, bool) {
	return Person{Name: name, Email: email}, true
}

func (p Person) GetJSON() (string, bool) {
	b, ok := json.Marshal(p)
	if ok != nil {
		return "", false
	}
	return string(b), true
}
