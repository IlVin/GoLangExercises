package main

import (
	"fmt"
	"math"
)

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

var area_calculators = make(map[figures]func(float64) float64)

func init() {
	area_calculators[square] = func(a float64) float64 {
		return a * a
	}
	area_calculators[circle] = func(r float64) float64 {
		return r * r * math.Pi
	}
	area_calculators[triangle] = func(a float64) float64 {
		p := (a + a + a) / 2
		return math.Sqrt(p * (p - a) * (p - a) * (p - a))
	}
}

func area(f figures) (func(float64) float64, bool) {
	fn, ok := area_calculators[f]
	if !ok {
		return nil, false
	}

	return fn, true
}

func main() {

	ar, ok := area(square)
	if !ok {
		fmt.Println("Ошибка")
		return
	}
	myArea := ar(5)

	fmt.Println(myArea)
}
