package main

import (
	"fmt"
	"strings"
)

func checkStr(s string) bool {
	// Преобразуем строку в нижнему регистру.
	lowerStr := strings.ToLower(s)
	// Карта для подсчета символов в строке.
	rMap := map[rune]int{}
	// Итерируемся по строке, каждый символ используем в качестве ключа в карте
	// и инкерментируем значение этого ключа.
	for _, r := range lowerStr {
		rMap[r]++
	}
	// Итерируемся по карте, если значение ключа больше 1, значит строке есть
	// повторяющийся символ, в этом случае возвращаем false.
	for _, v := range rMap {
		if v > 1 {
			return false
		}
	}
	// Если в карте значения всех ключей меньше 2, то в строке все
	// символы уникальны, возвращаем true.
	return true
}

func main() {
	fmt.Println(checkStr("abcd"))      // true
	fmt.Println(checkStr("abCdefAaf")) // false
	fmt.Println(checkStr("aabcd"))     // false
	fmt.Println(checkStr(""))          // true
}
