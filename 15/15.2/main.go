package main

import (
	"fmt"
	"unsafe"
)

var justString string

func createHugeString(l int) string {
	return string(make([]byte, l))
}

func someFunc() {
	v := createHugeString(1 << 10)
	// переводим строку в срез рун.
	runeSlice := []rune(v)
	// создаем новую строку в новой области памяти.
	justString = string(runeSlice[:100])
	fmt.Println(unsafe.StringData(v))
	fmt.Println(unsafe.StringData(justString))
}

func main() {
	someFunc()
}
