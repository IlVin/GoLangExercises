package main

import (
	"fmt"

	"github.com/IlVin/gomod/pkg/testmod"
)

func main() {
	if sum := testmod.Add(1, 2); sum != 3 {
		panic(fmt.Sprintf("sum expected to be 3; got %d", sum))
	}

	fmt.Println("Well done!")
}
