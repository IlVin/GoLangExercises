package main

import "fmt"

var Global = 5

func ChangeGlobalVar() {
	defer func(old_value int) {
		Global = old_value
	}(Global)
	Global = 15
	fmt.Println(Global)
}

func main() {

	fmt.Println("START")

	fmt.Println("GlobalVar START:", Global)
	ChangeGlobalVar()
	fmt.Println("GlobalVar STOP:", Global)
}
