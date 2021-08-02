package main

import "fmt"

func main() {
	var arr1 = []int{6, 23, 4, 1, 4, 35, 35, 35, 35, 7, 567, 8, 67}
	var arr2 = []int{35, 6, 4, 22}

	m1 := make(map[int]int)
	for _, v := range arr1 {
		m1[v] = 1
	}
	res := []int{}
	for _, v := range arr2 {
		if _, ok := m1[v]; ok {
			res = append(res, v)
		}
	}
	//O(len(arr1)+len(arr2))
	fmt.Println(res)
}
