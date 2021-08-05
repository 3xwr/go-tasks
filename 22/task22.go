package main

import "fmt"

func main() {
	arr := []int{9, 13, 4245, 53, 13, 57, 2, 3, 1, 0, -6, -80}
	fmt.Println(quickSort(arr))
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	//get left and right elements
	left, right := 0, len(arr)-1
	//pick the rightmost element as pivot
	pivot := right

	//displace all elements less than the pivot to the left
	for i := range arr {
		if arr[i] < arr[pivot] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	//move pivot in the middle
	arr[left], arr[pivot] = arr[pivot], arr[left]

	//recursively sort left and right parts
	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}
