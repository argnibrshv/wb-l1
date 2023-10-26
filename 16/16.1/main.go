package main

import "fmt"

func Partition(arr []int, leftIndex, rightIndex int) int {
	// определение опорного элемента pivot как среднего элемента в срезе arr.
	pivot := arr[(leftIndex + (rightIndex-leftIndex)/2)]
	for leftIndex <= rightIndex {
		// увеличиваем leftIndex на 1 пока элемент с индексом leftIndex
		// меньше опорного элемента pivot.
		for arr[leftIndex] < pivot {
			leftIndex++
		}
		// уменьшаем rightIndex на 1 пока опорный элемент pivot меньше
		// элемента с индексом rightIndex.
		for pivot < arr[rightIndex] {
			rightIndex--
		}
		if leftIndex <= rightIndex {
			// меняем местами значениями элементов с индексами leftIndex
			// и rightIndex в срезе arr.
			arr[leftIndex], arr[rightIndex] = arr[rightIndex], arr[leftIndex]
			leftIndex++
			rightIndex--
		}
	}
	// возвращаем значения leftIndex, которое будет использоваться как
	// индекс опорного элемента для дальнейшей сортировки.
	return leftIndex
}

func QuickSort(arr []int, leftBound, rightBound int) {
	if leftBound < rightBound {
		// определение опорного элемента pivot.
		pivot := Partition(arr, leftBound, rightBound)
		// рекурсивный вызов функции QuickSort для сортировки левой части
		// среза от leftBound до pivot-1.
		QuickSort(arr, leftBound, pivot-1)
		// рекурсивный вызов функции QuickSort для сортировки правой части
		// среза от pivot до rightBound.
		QuickSort(arr, pivot, rightBound)
	}
}

func main() {
	arr := []int{8, 2, 7, 3, 9, 1, 5, 0, 6, 4}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
