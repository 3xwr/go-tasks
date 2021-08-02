package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	concMap := make(map[int]int)
	mu := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go writeMap(concMap, mu)
	}
	readMap(concMap, mu)
}

func readMap(concMap map[int]int, mu *sync.Mutex) {
	mu.Lock()
	fmt.Println(concMap)
	mu.Unlock()
}

func writeMap(concMap map[int]int, mu *sync.Mutex) {
	mu.Lock()
	concMap[rand.Int()%100] = rand.Int() % 100
	mu.Unlock()
}
