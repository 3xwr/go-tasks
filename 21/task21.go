package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range arr {
		wg.Add(1)
		go Read(i, &wg)
	}
	wg.Wait()
}

func Read(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	wg.Done()
}
