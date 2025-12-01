package main

import "fmt"

func main() {
	var s = []int{}
	for i := 1; i <= 100; i++ {
		s = append(s, i)
	}

	s = append(s[:10], s[90:]...)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	for i, v := range s {
		fmt.Printf("s[%d] = %d\n", i, v)
	}

	input := "The quick brown 狐 jumped over the lazy 犬"

	// создаём слайс рун
	runes := make([]rune, 0, len(input))

	// добавляем руны в слайс
	for _, r := range input {
		runes = append(runes, r)
	}

	// разворачиваем
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// преобразуем в строку UTF-8.
	output := string(runes)
	fmt.Println(output)

	m := map[rune]int{'a': 1, 'b': 3}
	fmt.Println(m)
	delete(m, 'a')
	fmt.Println("DEL: ", m)
}
