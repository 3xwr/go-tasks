package main

import (
	"fmt"
	"time"
)

func printUnbuf(ch chan int) {
	fmt.Println(<-ch)
	fmt.Println("Hi from unbuf!")
}

func printBuf(ch chan int) {
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("Hi from buf!")
}

func main() {
	ch := make(chan int)
	chBuf := make(chan int, 3)
	go printUnbuf(ch)
	go printBuf(chBuf)
	ch <- 0
	//ch <- 1 Вызовет дедлок, потому что мейн горутина будет ждать чтения из канала.
	chBuf <- 1
	chBuf <- 2
	chBuf <- 3 //не вызывает дедлок, т.к. printBuf прочитала два значения и в буфере одно значение,
	//горутина не блокируется
	time.Sleep(time.Second)
}
