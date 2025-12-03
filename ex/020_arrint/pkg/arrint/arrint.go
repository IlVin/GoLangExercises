package arrint

import (
	"fmt"
	"strings"
)

// Тип: Слайс чиселок
type ArrInt []int

// Возвращает слайс, значения элементов которого равны сумме значений элементов входящих слайсов
// Длина результирующего массива равна наименьшей длине входящих массивов
func Add(a, b ArrInt) ArrInt {
	length := len(a)
	if length-len(b) > 0 {
		length = len(b)
	}
	c := make(ArrInt, length)
	for i := 0; i < length; i++ {
		c[i] = a[i] + b[i]
	}
	return c
}

// Метод String преобразует ArrInt в строку и возвращает её.
func (a ArrInt) String() string {
	out := []string{}

	for _, v := range a {
		out = append(out, fmt.Sprintf(`<%d>`, v))
	}
	return fmt.Sprintf(`[%s]`, strings.Join(out, ` `))
}
