package main

import (
	"fmt"
	"time"
)

func SendInChannel(out chan<- int, N int) {
	// Создаем канал, в который по истечению таймера придет значение времени.
	timeout := time.After(time.Duration(N) * time.Second)
	// Значение, которое отправляем в канал.
	value := 0
	for {
		// Будет производиться запись в канал out, пока не истечет таймер, по истечению таймера закроется канал out, будет произведен выход из функции.
		select {
		case out <- value:
			value++
		case <-timeout:
			close(out)
			return
		}
	}
}

func main() {
	ch := make(chan int)
	// Запускает в горутине функцию отправки в канал и передаем ей
	// канал и количество секунд, по истечению которых программа
	// должна завершиться.
	go SendInChannel(ch, 1)
	// Последовательно читаем из канала, пока он не будет закрыт,
	// по истечению таймера, заданного в функции SendInChannel.
	for value := range ch {
		fmt.Println(value)
	}
}
