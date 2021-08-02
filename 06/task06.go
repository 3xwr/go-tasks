package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool)
	go example(quit)
	time.Sleep(2 * time.Second)
	quit <- true

	ctx := context.Background()
	go example2(ctx)
	time.Sleep(2 * time.Second)
	ctx.Done()
}

func example(quit chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("Second spent on example #1")
		}
	}
}

func example2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("Second spent on example #2")
		}
	}
}
