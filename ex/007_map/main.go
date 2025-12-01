package main

import "fmt"

func main() {

	price := map[string]float64{"хлеб": 50, "молоко": 100, "масло": 200, "колбаса": 500, "соль": 20, "огурцы": 200, "сыр": 600, "ветчина": 700, "буженина": 900, "помидоры": 250, "рыба": 300, "хамон": 1500.}

	fmt.Println("Деликатесы:")

	for k, v := range price {
		if v > 500 {
			fmt.Printf("%s - %0.2f\n", k, v)
		}
	}

	order := []string{"хлеб", "буженина", "сыр", "огурцы"}

	var total float64
	for _, v := range order {
		total += price[v]
	}

	fmt.Println("TOTAL: ", total)
}
