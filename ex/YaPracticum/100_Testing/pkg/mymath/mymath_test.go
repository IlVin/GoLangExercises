package mymath

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func IntsToString(a ...int) string {
	res := strings.Builder{}
	for _, v := range a {
		res.WriteString(fmt.Sprintf("%d, ", v))
	}
	return res.String()
}

func FloatsToString(f ...float64) string {
	res := strings.Builder{}
	for _, v := range f {
		res.WriteString(fmt.Sprintf("%f, ", v))
	}
	return res.String()
}

func TestAddInts(t *testing.T) {
	TestPlan := []struct {
		name string
		args []int
		got  int
	}{
		{name: "Empty args", args: []int{}, got: 0},
		{name: "One arg", args: []int{5}, got: 5},
		{name: "Two args", args: []int{7, 8}, got: 15},
		{name: "Many non negative args", args: []int{7, 8, 0, 45, 67}, got: 127},
		{name: "One negative arg", args: []int{-5}, got: -5},
		{name: "Two negative args", args: []int{-7, -8}, got: -15},
		{name: "Meny non positive args", args: []int{-7, -8, 0, -45, -67}, got: -127},
	}
	{
		for _, plan := range TestPlan {
			t.Run(plan.name, func(t *testing.T) {
				if sum := AddInts(plan.args...); sum != plan.got {
					t.Errorf("TestAddInts(%s)=%d, got %d", IntsToString(plan.args...), sum, plan.got)
				}
			})
		}
	}
}

func TestAbs(t *testing.T) {
	TestPlan := []struct {
		name string
		arg  float64
		got  float64
	}{
		{name: "Null", arg: 0.0, got: 0},
		{name: "Abs(-0.1)", arg: -0.1, got: 0.1},
		{name: "Abs(0.1)", arg: 0.1, got: 0.1},
		{name: "Abs(-0.5)", arg: -0.5, got: 0.5},
		{name: "Abs(0.5)", arg: 0.5, got: 0.5},
		{name: "Abs(-0.6)", arg: -0.6, got: 0.6},
		{name: "Abs(0.6)", arg: 0.6, got: 0.6},

		{name: "Abs(-3)", arg: -3, got: 3},
		{name: "Abs(3)", arg: 3, got: 3},
		{name: "Abs(-2.000001)", arg: -2.000001, got: 2.000001},
		{name: "Abs(-0.000000003)", arg: -0.000000003, got: 0.000000003},
	}
	{
		for _, plan := range TestPlan {
			t.Run(plan.name, func(t *testing.T) {
				if res := Abs(plan.arg); !(res >= 0 && Abs(res-plan.got) <= math.SmallestNonzeroFloat64) {
					t.Errorf("Abs(%s)=%f, got %f", FloatsToString(plan.arg), res, plan.got)
				}
			})
		}
	}
}
