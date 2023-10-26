package main

import (
	"fmt"
)

// changeBit по индексу изменяет бит на 0 или 1.
func changeBit(num int64, idx int) int64 {
	if idx < 0 || idx > 63 {
		fmt.Println("Неверный номер бита")
		return num
	}
	if idx == 63 {
		// 64 бит обозначает знак числа, возвращая противоположное значение числа,
		// возвращается число с измененным битом знака числа.
		return -num
	} else {
		// создаем битовую маску, сдвигая 1 на необходимое количество разрядов,
		// и применяем побитовое исключающее ИЛИ.
		return num ^ (1 << idx)
	}
}

func main() {
	var example int64 = 2456
	res := changeBit(example, 63)
	fmt.Printf("%064b\n%064b\n%d\n", example, res, res)
}
