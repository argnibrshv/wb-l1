package main

import (
	"fmt"
	"reflect"
)

func DetermineType(v interface{}) {
	// Метод TypeOf из пакета reflect, возвращает тип переменной переданный в него.
	fmt.Println("type is", reflect.TypeOf(v))
}

func main() {
	var test interface{}
	test = "string"
	DetermineType(test) // type is string
	test = 22
	DetermineType(test) // type is int
	test = true
	DetermineType(test) // type is bool
	test = make(chan int)
	DetermineType(test) // type is chan int
	test = make(chan interface{})
	DetermineType(test) // type is chan interface {}
}
