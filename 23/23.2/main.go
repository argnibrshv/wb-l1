package main

import "fmt"

// удаление элемента из среза без сохранения порядка элементов
func removeElement(slc []int, idx int) ([]int, error) {
	// если индекс удаляемого элемента выходит за пределы среза,
	// возвращаем исходный срез и ошибку.
	if idx >= len(slc) || idx < 0 {
		return slc, fmt.Errorf("index is out of range. index is %d, slice length %d", idx, len(slc))
	}
	lastIdx := len(slc) - 1
	// копируем последний элемент в индекс удаляемого элемента
	slc[idx] = slc[lastIdx]
	// возвращаем срез без послежнего элемента
	return slc[:lastIdx], nil
}

func main() {
	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	testSlice, err := removeElement(testSlice, 0)
	checkError(err)
	fmt.Println(testSlice) // [8 2 3 4 5 6 7]
	testSlice, err = removeElement(testSlice, 6)
	checkError(err)
	fmt.Println(testSlice) // [8 2 3 4 5 6]
	testSlice, err = removeElement(testSlice, 3)
	checkError(err)
	fmt.Println(testSlice) // [8 2 3 6 5]
	testSlice, err = removeElement(testSlice, -1)
	checkError(err)        // index is out of range. index is -1, slice length 5
	fmt.Println(testSlice) // [8 2 3 6 5]
	testSlice, err = removeElement(testSlice, 6)
	checkError(err)        // index is out of range. index is 6, slice length 5
	fmt.Println(testSlice) // [8 2 3 6 5]
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
