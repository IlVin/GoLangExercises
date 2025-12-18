package myuser

import (
	"testing"
)

func TestFullName(t *testing.T) {

	test_plan := []struct {
		name string
		u    User
		got  string
	}{
		{name: "Full data", u: User{FirstName: "Misha", LastName: "Popov"}, got: "Misha Popov"},
		{name: "Empty data", u: User{FirstName: "", LastName: ""}, got: ""},
		{name: "Default init", u: User{}, got: ""},
		{name: "Empty LastName", u: User{FirstName: "Misha", LastName: ""}, got: "Misha"},
		{name: "Empty FirstName", u: User{FirstName: "", LastName: "Popov"}, got: "Popov"},
	}

	for _, plan := range test_plan {
		t.Run(plan.name, func(t *testing.T) {
			if res := plan.u.FullName(); plan.got != res {
				t.Errorf("Result='%s', got='%s'", res, plan.got)
			}
		})
	}
}
