package main

import "fmt"

// Последовательно полученные данные из канала in удваиваются и отправляются в канал out, когда канал in будет закрыт, закрываем канал out.
func Double(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * 2
	}
	close(out)
}

// Функция последовательно получает данные из канала и выводит их в stdout,
//
//	когда канал будет закрыт, функция завершиться
func PrintResult(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	// Канал для передачи чисел из массива.
	ch1 := make(chan int)
	// Канал для передачи удвоенного значения.
	ch2 := make(chan int)
	arr := [6]int{2, 4, 6, 8, 10, 12}
	// Функция в горутине принимает канал для передачи чисел из массива в
	// качестве аргумента, каждое число из массива последовательно отправляем в
	// канал, после чего канал закрывается.
	go func(out chan<- int) {
		for _, num := range arr {
			out <- num
		}
		close(out)
	}(ch1)
	// Запускаем горутину с функцией для удвоения чисел.
	go Double(ch1, ch2)
	PrintResult(ch2)
}
