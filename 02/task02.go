package main

import (
	"fmt"
	"sync"
)

func square(num int, channel chan int) {
	channel <- num * num
}

func square_wg(wg *sync.WaitGroup, num int) {
	fmt.Println(num * num)
	wg.Done()
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	output := make(chan int, cap(nums))
	//way 1
	fmt.Println("Square using channels: ")
	for _, v := range nums {
		go square(v, output)
	}
	for i := 0; i < cap(nums); i++ {
		fmt.Println(<-output)
	}

	//way 2
	fmt.Println("Square using waitGroup: ")
	var wg sync.WaitGroup
	for _, v := range nums {
		wg.Add(1)
		go square_wg(&wg, v)
	}
	wg.Wait()
}
