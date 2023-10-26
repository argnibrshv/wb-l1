package main

import (
	"fmt"
	"math"
	"os"
	"sync"
)

func squareOfNum(num int, wg *sync.WaitGroup) {
	// отложенный вызов метода Done для объекта WaitGroup.
	defer wg.Done()
	// вычисление квадрата числа и присвоение результата переменной result.
	result := math.Pow(float64(num), 2)
	// вывод результата в стандартный поток вывода с помощью метода
	// Fprintln из пакета fmt.
	fmt.Fprintln(os.Stdout, result) // fmt.Println(result) будет так же выводить в Stdout
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	// создание объекта WaitGroup для синхронизации горутин.
	var wg sync.WaitGroup
	// цикл перебора элементов массива arr.
	for _, num := range arr {
		// добавление одного ожидаемого завершения горутины к объекту WaitGroup.
		wg.Add(1)
		// запуск горутины для вычисления квадрата числа.
		go squareOfNum(num, &wg)
	}
	// ожидание завершения всех горутин с помощью метода Wait объекта WaitGroup.
	wg.Wait()
}
