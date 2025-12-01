package main

import (
	"fmt"
	"strings"
)

func CntCaller(f func(string) string) (wrapper func(string) string, result func() int) {
	cnt := 0
	wrapper = func(s string) string {
		cnt++
		return f(s)
	}
	result = func() int {
		return cnt
	}
	return
}

func OrigFunc(s string) string {
	fmt.Println("Call original func: ", s)
	return strings.Join([]string{"called orig func", s}, ":")
}

func main() {
	f, res := CntCaller(OrigFunc)

	fmt.Println(f("+1"), "CNT: ", res())
	fmt.Println(f("+3"), "CNT: ", res())
	fmt.Println(f("+5"), "CNT: ", res())
}
