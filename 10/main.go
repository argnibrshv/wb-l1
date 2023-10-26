package main

import "fmt"

func main() {
	// последовательность температурных колебаний
	var t = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// Создаем карту, где будут храниться группы температур
	tMap := make(map[int][]float64)
	// итерируемся по всей последовательности температурных колебаний
	for _, v := range t {
		// пустой оператор switch позволяет использовать для каждой ветви case
		// любую операцию сравнения, дающую в результате логическое значение.
		switch {
		case v > -30 && v <= -20:
			tMap[-20] = append(tMap[-20], v)
		case v >= 10 && v < 20:
			tMap[10] = append(tMap[10], v)
		case v >= 20 && v < 30:
			tMap[20] = append(tMap[20], v)
		case v >= 30 && v < 40:
			tMap[30] = append(tMap[30], v)
		}
	}
	for k, v := range tMap {
		fmt.Printf("%v:%v\n", k, v)
	}
}
