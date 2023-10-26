package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseString(s string) string {
	// Переменная равная количеству символов в исходной строку
	// для последовательного изменений значений под этим индексом в
	// срезе рун для развернутой строки.
	n := utf8.RuneCountInString(s)
	// срез рун, для хранения символов развернутой строки
	runes := make([]rune, n)
	// Итерируемся по строке и каждую руну этой строки назначаем
	// в срез рун в обратном порядке с помощью переменной n,
	// которую уменьшаем на единицу при каждой итерации.
	for _, r := range s {
		n--
		runes[n] = r
	}
	return string(runes)
}

func main() {
	testStr := "главрыба"
	reversedStr := reverseString(testStr)
	fmt.Println(reversedStr) // абырвалг
	testStr = "猫的妻子"
	reversedStr = reverseString(testStr)
	fmt.Println(reversedStr) // 子妻的猫
	fmt.Println(reverseString("test"))
}
