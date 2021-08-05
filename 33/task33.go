package main

import (
	"fmt"
	"math/rand"
	"time"
)

func checkEven(ch1 chan int, ch2 chan int) {
	for {
		num := <-ch1
		if num%2 == 0 {
			ch2 <- num
		}
	}
}

func printNums(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ch1 := make(chan int)
	ch2 := make(chan int)
	go checkEven(ch1, ch2)
	go printNums(ch2)

	for i := 0; i < 10; i++ {
		num := rand.Int() % 100
		ch1 <- num
	}
}
