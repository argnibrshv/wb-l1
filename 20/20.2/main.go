package main

import (
	"fmt"
	"strings"
)

func reverseWord(s string) string {
	// Получаем срез строк из слов, разделенных пробелом.
	words := strings.Split(s, " ")
	n := len(words)
	// Срез для хранения слов в обратном порядке
	reverseSlice := make([]string, n)
	// Итерируемся по срезу words и каждую строку этого стреза назначаем
	// в срез reverseSlice в обратном порядке с помощью переменной n,
	// которую уменьшаем на единицу при каждой итерации.
	for _, word := range words {
		n--
		reverseSlice[n] = word
	}
	return strings.Join(reverseSlice, " ")
}

func main() {
	fmt.Println(reverseWord("snow dog sun run man can want"))
}
