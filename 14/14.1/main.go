package main

import (
	"fmt"
)

// Оператор switch использует специальное утверждение типа, в котором
// используется ключевое слово type. Каждый оператор case указывает тип
// и блок кода, который будет выполняться, когда значение, оцениваемое
// оператором switch, имеет указанный тип. Оператор default используется для
// указания блока кода, который будет выполняться, когда ни один
// из операторов case не cовпадает.
func DetermineType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
	case bool:
		fmt.Println("type is bool")
	default:
		// При форматировании строк с помощью printf используя глагол %T, выводиться тип значения.
		fmt.Printf("type is %T\n", v)
	}
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
