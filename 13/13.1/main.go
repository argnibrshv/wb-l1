package main

// Swap меняет местами значения элементов с индексами i1 и i2 в срезе arr.
// Множественное присваивание позволяет обменять значения между переменными
// без необходимости создавать временные переменные.
func Swap(slc []int, i1, i2 int) {
	slc[i1], slc[i2] = slc[i2], slc[i1]
}

func main() {
	slc := []int{1, 2, 3, 4}
	Swap(slc, 0, 3) // [4, 2, 3, 1]
	Swap(slc, 1, 3) // [4, 1, 3, 2]
}
