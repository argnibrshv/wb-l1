package main

import (
	"fmt"
	"sync"
	"time"
)

func changeMap(m map[int]int, mu *sync.Mutex, key int) {
	// Устанавливаем блокировку мьютекса, доступ к карте будет
	// только у одной горутины.
	mu.Lock()
	// Снимаем блокировку мьютекса после завершения функции.
	defer mu.Unlock()
	// Изменяем карту.
	m[key]++
}

func main() {
	// Создаем карту, которую будем производить конкурентную запись данных.
	countMap := make(map[int]int)
	// Определяем мьютекс.
	var mu sync.Mutex
	// Запускаем 10 горутин для конкурентной записи в карту.
	for i := 0; i < 10; i++ {
		go changeMap(countMap, &mu, i)
	}
	// Ждем секунду,чтобы другие горутины завершились.
	time.Sleep(time.Second)
	fmt.Println(countMap)
}
