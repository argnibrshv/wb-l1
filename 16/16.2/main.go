package main

import "fmt"

func QuickSort(arr []int) []int {
	// условие, которое проверяет, что длина среза меньше 2,
	// чтобы избежать бесконечной рекурсии.
	if len(arr) < 2 {
		return arr
	} else {
		// определение опорного элемента pivot как середины среза arr.
		pivot := len(arr) / 2
		// создание двух новых срезов left и right для хранения элементов
		// меньше и больше опорного элемента соответственно.
		left := []int{}
		right := []int{}
		for _, v := range arr {
			// Элементы среза меньше опорного добавляются в левый срез,
			// элементы больше опорного добавляются в правый срез.
			if v < arr[pivot] {
				left = append(left, v)
			} else if v > arr[pivot] {
				right = append(right, v)
			}
		}
		// рекурсивный вызов функции QuickSort для сортировки левой части
		// среза от начала до опорного элемента и добавление опорного элемента
		// в конец отсортированного среза.
		left = append(QuickSort(left), arr[pivot])
		// рекурсивный вызов функции QuickSort для сортировки правой части
		// среза от опорного элемента до конца и объединение срезов left
		// и отсортированного right.
		return append(left, QuickSort(right)...)
	}

}

func main() {
	arr := []int{8, 2, 7, 3, 9, 1, 5, 0, 6, 4}
	fmt.Println(QuickSort(arr))
}
