package main

import "fmt"

func main() {
	// первое неупорядоченное множество.
	set1 := map[int]bool{
		5: true,
		2: true,
		9: true,
		7: true,
	}
	// второе неупорядоченное множество.
	set2 := map[int]bool{
		5: true,
		8: true,
		4: true,
		2: true,
	}
	// пересечение множеств.
	intersection := make(map[int]bool)
	// карта в которой будет храниться количество вхождений в каждое множество.
	valueSetsMap := make(map[int]int)
	// итерируемся по множествам и добавляем в карту количество вхождений
	// элементов во множествах.
	for k := range set1 {
		valueSetsMap[k]++
	}
	for k := range set2 {
		valueSetsMap[k]++
	}
	// итерируемся по карте, ключи со значение 2 встречаются в обоих множествах,
	// добавялем эти ключи в пересечение множеств.
	for k, v := range valueSetsMap {
		if v == 2 {
			intersection[k] = true
		}
	}
	// выводим пересечения множеств в консоль.
	for k := range intersection {
		fmt.Printf("%d ", k)
	}
	fmt.Printf("\n")
}
