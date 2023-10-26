package main

import "fmt"

func binarySearch(arr []int, lowIdx, highIdx, target int) int {
	// базовый случай рекурсивной функции.
	if highIdx < lowIdx {
		return -1
	} else {
		// находим индекс середины массива.
		mid := lowIdx + (highIdx-lowIdx)/2
		// если искомое значение больше значения индекса середины массива,
		// то рекурсивно вызываем функцию бинарного поиска для правой
		// половины среза.
		// если искомое значение меньше значения индекса середины массива,
		// то рекурсивно вызываем функцию бинарного поиска для левой
		// половины среза.
		// если искомое значение ни больше, ни меньше значения индекса
		// середины массива, значит значения индекса середины массива и
		// есть искомое значение.
		if arr[mid] < target {
			return binarySearch(arr, mid+1, highIdx, target)
		} else if arr[mid] > target {
			return binarySearch(arr, lowIdx, mid-1, target)
		} else {
			return mid
		}
	}
}

// Обертка для удобного запуска бинарного поиска
func binarySearchStart(arr []int, target int) int {
	return binarySearch(arr, 0, len(arr)-1, target)
}

func main() {
	testArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(binarySearchStart(testArr, 5))  //4
	fmt.Println(binarySearchStart(testArr, 1))  //0
	fmt.Println(binarySearchStart(testArr, 0))  //-1
	fmt.Println(binarySearchStart(testArr, 11)) //10
	fmt.Println(binarySearchStart(testArr, 12)) //-1
	fmt.Println(binarySearchStart(testArr, 22)) //-1
}
