package main

import (
	"fmt"
)

func main() {

	m := [3]int{1, 2, 3}

	for i, v := range &m {
		m[i] = 4
		fmt.Println(v)
	}

	fmt.Println(m)
}
