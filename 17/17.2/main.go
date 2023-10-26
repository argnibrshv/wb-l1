package main

import "fmt"

func binarySearch(arr []int, target int) int {
	// нижняя граница поиска
	lowIdx := 0
	// верхняя граница поиска
	highIdx := len(arr) - 1
	for lowIdx <= highIdx {
		// находим индекс середины массива
		mid := lowIdx + (highIdx-lowIdx)/2
		// если искомое значение больше значения индекса середины массива,
		// то изменяем нижнюю границу поиска на индекс середины массива
		// увеличенный на единицу.
		// если искомое значение меньше значения индекса середины массива, то
		// изменяем верхнюю границу поиска на индекс середины массива
		// уменьшенный на единицу.
		// если искомое значение равно значению индекса середины массива,
		// возвращаем индекс середины массива.
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			lowIdx = mid + 1
		} else if arr[mid] > target {
			highIdx = mid - 1
		}
	}
	// если в цикле не найден искомый элемент, значит его нет в массиве,
	// возвращаем минус единицу
	return -1
}

func main() {
	testArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(binarySearch(testArr, 5))  //4
	fmt.Println(binarySearch(testArr, 1))  //0
	fmt.Println(binarySearch(testArr, 0))  //-1
	fmt.Println(binarySearch(testArr, 11)) //10
	fmt.Println(binarySearch(testArr, 12)) //-1
	fmt.Println(binarySearch(testArr, 22)) //-1
}
