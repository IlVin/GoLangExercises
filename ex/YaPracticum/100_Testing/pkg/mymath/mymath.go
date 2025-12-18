package mymath

import "math"

func AddInts(args ...int) int {
	sum := 0
	for _, num := range args {
		sum += num
	}
	return sum
}

// Abs возвращает абсолютное значение.
// Например: 3.1 => 3.1, -3.14 => 3.14, -0 => 0.
func Abs(value float64) float64 {
	return math.Abs(value)
}
