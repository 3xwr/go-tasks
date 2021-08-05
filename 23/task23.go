package main

import "fmt"

func main() {
	arr := []int{1, 12, 23, 44, 67, 87, 100}
	searchItem := 44
	fmt.Printf("Item %d found at index %d using iterative binary search.\n",
		searchItem, binarySearch(arr, searchItem))
	fmt.Printf("Item %d found at index %d using recursive binary search.\n",
		searchItem, binarySearchRecursive(arr, 0, len(arr), searchItem))
}

//iterative
func binarySearch(arr []int, item int) int {
	if len(arr) < 2 {
		return 0
	}
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] < item {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if arr[left] != item {
		return -1
	}

	return left
}

func binarySearchRecursive(arr []int, left int, right int, item int) int {
	if len(arr) < 2 {
		return 0
	}
	mid := (left + right) / 2
	if item < arr[mid] {
		return binarySearchRecursive(arr, left, mid-1, item)
	} else if item > arr[mid] {
		return binarySearchRecursive(arr, mid+1, right, item)
	} else if item == arr[mid] {
		return mid
	}
	return -1
}
