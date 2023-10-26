package main

import (
	"fmt"
	"strings"
)

func reverseWord(s string) string {
	// Получаем срез строк из слов, разделенных пробелом.
	words := strings.Split(s, " ")
	// Для цикла Создаем переменную left для индекса начала среза
	// и переменную right для индекса конца среза, цикл будет
	// продолжаться пока значение left меньшe right,
	// на каждой итерации меняем места в срезе значения под индексами
	// left и right, инкрементируем занчение left и декрементируем значение right.
	for left, right := 0, len(words)-1; left < right; left++ {
		words[left], words[right] = words[right], words[left]
		right--
	}
	// Объеденяем срез в строку и возвращаем ее.
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWord("snow dog run sun"))
}
