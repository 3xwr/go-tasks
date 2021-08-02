package main

import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go in_func(in, out)
	go out_func(out)
	for _, v := range arr {
		in <- v
	}
}

func in_func(in chan int, out chan int) {
	for {
		out <- 2 * <-in
	}
}

func out_func(out chan int) {
	for {
		fmt.Println(<-out)
	}
}
