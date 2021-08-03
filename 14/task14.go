package main

import "fmt"

type Set struct {
	items []string
}

func (s *Set) makeSet(arr []string) {
	tempMap := make(map[string]int)
	for _, v := range arr {
		tempMap[v] = 1
	}
	for k := range tempMap {
		s.items = append(s.items, k)
	}
}

func main() {
	var resSet Set
	var arr = []string{"cat", "cat", "dog", "cat", "tree"}
	resSet.makeSet(arr)
	fmt.Println(resSet)
}
