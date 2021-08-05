package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	num int
}

func (c *Counter) increase() {
	c.Lock()
	c.num++
	c.Unlock()
}

func (c *Counter) printCount() {
	c.Lock()
	fmt.Println(c.num)
	c.Unlock()
}

func main() {
	counter := Counter{
		num: 0,
	}
	for i := 0; i < 5; i++ {
		go counter.increase()
	}
	time.Sleep(time.Nanosecond * 50)
	counter.printCount()
}
