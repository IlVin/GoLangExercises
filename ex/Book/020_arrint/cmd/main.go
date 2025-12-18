package main

import (
	"arrint/pkg/arrint"
	"fmt"
)

func main() {
	fmt.Println("arrint")
	fmt.Printf("%q\n", arrint.Add([]int{1, 12, 3}, []int{9, 7}).String())
}
