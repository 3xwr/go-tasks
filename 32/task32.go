package main

import (
	"fmt"
	"time"
)

func mySleep(seconds int) {
	<-time.After(time.Second * time.Duration(seconds))
}

func main() {
	fmt.Println("Hello")
	mySleep(3)
	fmt.Println("World!")
}
