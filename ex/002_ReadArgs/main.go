package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Вывод аргументов коммандной строки
	fmt.Println("Args: " + strings.Join(os.Args, ", "))

	// Счетчик уникальных строк
	var cnt = make(map[string]int)

	// Сканнер STDIN
	var reader = bufio.NewScanner(os.Stdin)

	// Чтение STDIN и подсчет уникальных строк
	for reader.Scan() {
		cnt[reader.Text()]++
		fmt.Println(reader.Text())
	}

	// Вывод того, что насчитали
	fmt.Println(cnt)
}
