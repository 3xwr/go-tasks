package main

import "fmt"

func deleteFromSliceByIndex(slice []int, i int) []int {
	//... передает все элементы от i+1 включительно
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice = deleteFromSliceByIndex(slice, 0)
	slice = deleteFromSliceByIndex(slice, 6)
	fmt.Println(slice)
}
