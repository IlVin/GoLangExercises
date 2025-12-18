package myfamily

import "testing"

type PersonList []struct {
	person Person
	role   Relationship
}

type Plan struct {
	name    string
	persons PersonList
	got     error
}

func TestAddNew(t *testing.T) {
	test_plan := []Plan{
		{
			name: "No dup",
			persons: PersonList{
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: Father},
				{person: Person{FirstName: "Mother", LastName: "Popova", Age: 2}, role: Mother},
				{person: Person{FirstName: "Child", LastName: "Popov", Age: 2}, role: Child},
				{person: Person{FirstName: "GrandMother", LastName: "Popova", Age: 2}, role: GrandMother},
				{person: Person{FirstName: "GrandFather", LastName: "Popov", Age: 2}, role: GrandFather},
			},
			got: nil,
		},
		{
			name: "Traditional complete family with many children",
			persons: PersonList{
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: Father},
				{person: Person{FirstName: "Mother", LastName: "Popova", Age: 2}, role: Mother},
				{person: Person{FirstName: "Son1", LastName: "Popov", Age: 2}, role: Child},
				{person: Person{FirstName: "Daughter2", LastName: "Popov", Age: 2}, role: Child},
				{person: Person{FirstName: "Son3", LastName: "Popov", Age: 2}, role: Child},
				{person: Person{FirstName: "GrandMother1", LastName: "Popova", Age: 2}, role: GrandMother},
				{person: Person{FirstName: "GrandMother2", LastName: "Popova", Age: 2}, role: GrandMother},
				{person: Person{FirstName: "GrandFather1", LastName: "Popov", Age: 2}, role: GrandFather},
				{person: Person{FirstName: "GrandFather2", LastName: "Popov", Age: 2}, role: GrandFather},
			},
			got: nil,
		},
		{
			name: "Two Father",
			persons: PersonList{
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: Father},
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: Father},
			},
			got: ErrRelationshipAlreadyExists,
		},
		{
			name: "Two Mother",
			persons: PersonList{
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: Mother},
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: Mother},
			},
			got: ErrRelationshipAlreadyExists,
		},
		{
			name: "Two Child",
			persons: PersonList{
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: Child},
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: Child},
			},
			got: nil,
		},
		{
			name: "Two GrandMother",
			persons: PersonList{
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: GrandMother},
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: GrandMother},
			},
			got: nil,
		},
		{
			name: "Two GrandFather",
			persons: PersonList{
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: GrandFather},
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: GrandFather},
			},
			got: nil,
		},
		{
			name: "3 GrandMother",
			persons: PersonList{
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: GrandMother},
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: GrandMother},
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: GrandMother},
			},
			got: ErrRelationshipAlreadyExists,
		},
		{
			name: "3 GrandFather",
			persons: PersonList{
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: GrandFather},
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: GrandFather},
				{person: Person{FirstName: "Misha", LastName: "Popov", Age: 2}, role: GrandFather},
			},
			got: ErrRelationshipAlreadyExists,
		},
	}

	for _, plan := range test_plan {
		t.Run(plan.name, func(t *testing.T) {
			f := Family{}
			var err error
			for _, p := range plan.persons {
				if e := f.AddNew(p.role, p.person); e != nil {
					err = e
				}
			}
			if err != plan.got {
				t.Errorf("ERR: %v", err)
			}
		})
	}
}
