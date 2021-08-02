package main

import (
	"fmt"
)

func square(num int, channel chan int) {
	channel <- num * num
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	output := make(chan int, cap(nums))
	total := 0
	//way 1
	fmt.Println("Sum of squares using channels: ")
	for _, v := range nums {
		go square(v, output)
	}
	for i := 0; i < cap(nums); i++ {
		total += <-output
	}
	fmt.Println(total)
}
