package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go send(ch)
	go receive(ch)
	time.Sleep(3 * time.Second)
}

func send(ch chan int) {
	for {
		ch <- rand.Int() % 10
	}
}

func receive(ch chan int) {
	for {
		fmt.Println(<-ch % 10)
	}
}
