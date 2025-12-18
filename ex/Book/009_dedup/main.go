package main

import "fmt"

func RemoveDuplicates(A []string) []string {
	s2b := map[string]bool{}
	r := []string{}
	for _, v := range A {
		if _, ok := s2b[v]; ok {
			continue
		}
		s2b[v] = true
		r = append(r, v)
	}
	return r
}

func main() {
	input := []string{"cat", "dog", "bird", "dog", "parrot", "cat"}

	input = RemoveDuplicates(input)

	fmt.Println("Dedup strings: ", input)

}
