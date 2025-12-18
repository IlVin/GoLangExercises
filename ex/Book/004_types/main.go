package main

import (
	"flag"
	"fmt"
)

func HelloGen(age int) string {
	switch {
	case age >= 1946 && age <= 1964:
		return "Привет, бумер!"
	case age >= 1965 && age <= 1980:
		return "Привет, представитель X!"
	case age >= 1981 && age <= 1996:
		return "Привет, миллениал!"
	case age >= 1997 && age <= 2012:
		return "Привет, зумер!"
	case age >= 2012:
		return "Привет, альфа!"
	}
	return "Привет, Deadley!"
}

func main() {
	age := flag.Int("age", 1900, "Age")
	flag.Parse()

	fmt.Println(HelloGen(*age))

	for i := 1; i < 100; i++ {
		switch {
		case i%5 == 0 && i%3 == 0:
			fmt.Println(i, ": ", "FizBuzz")
		case i%3 == 0:
			fmt.Println(i, ": ", "Fiz")
		case i%5 == 0:
			fmt.Println(i, ": ", "Buzz")
		default:
			fmt.Println(i)
		}
	}
	fmt.Println("End")

	a := 1
	p := &a
	b := &p

	*p = 3
	**b = 5

	a += 4 + *p + **b

	fmt.Printf("a = %d", a)
	fmt.Printf("*p = %d", *p)
	fmt.Printf("**b = %d", **b)
}
