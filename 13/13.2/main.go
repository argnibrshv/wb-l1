package main

import "fmt"

// Множественное присваивание позволяет обменять значения между переменными
// без необходимости создавать временные переменные.
func main() {
	one := 1
	two := 2
	one, two = two, one
	fmt.Println(one) // 2
	fmt.Println(two) // 1
}
