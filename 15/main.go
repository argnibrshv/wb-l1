package main

import (
	"fmt"
	"strings"
	"unsafe"
)

var justString string

func createHugeString(l int) string {
	return string(make([]byte, l))
}

func someFunc() {
	v := createHugeString(1 << 10)
	// Clone возвращает новую копию строки в новой области памяти.
	justString = strings.Clone(v[:100])
	// StringData возвращает указатель на базовые байты строки,
	// у justString и v они будут различаться.
	fmt.Println(unsafe.StringData(v))
	fmt.Println(unsafe.StringData(justString))
}

func main() {
	someFunc()
}
