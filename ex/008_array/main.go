package main

import "fmt"

func find_two_indexes(A []int, k int) []int {
	if A == nil {
		return nil
	}

	v2i := map[int]int{} // Value to index

	for i, v := range A {
		// Exists value k-v ?
		if v, ok := v2i[k-v]; ok {
			return []int{v2i[k-v], i}
		}
		v2i[v] = i
	}

	return nil
}

func main() {

	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 13, 154, 53, 2345, 5}
	k := 9000

	r := find_two_indexes(A[:], k)

	fmt.Println("Array: ", A)
	fmt.Println("k: ", k)
	fmt.Println("r: ", r)
}
