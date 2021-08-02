package main

import (
	"fmt"
	"time"
)

var N int

type Worker struct {
	ID   int
	ch   chan int
	done chan bool
}

func Work(num int) {
	fmt.Println(num)
}

func makeWorker(ch chan int, ID int) *Worker {
	TempWorker := &Worker{
		ID:   ID,
		ch:   ch,
		done: make(chan bool),
	}
	return TempWorker
}

func (w *Worker) Start() {
	fmt.Println("Starting worker #", w.ID)
	for {
		select {
		case i := <-w.ch:
			Work(i)
		case <-w.done:
			return
		}
	}
}

func (w *Worker) Stop() {
	w.done <- true
	fmt.Println("Stopping worker #", w.ID)
}

func main() {
	//N = num of workers
	N = 10
	data := []int{}

	//Data for workers
	for i := 0; i < N*10; i++ {
		data = append(data, i)
	}

	ch := make(chan int)
	workers := []Worker{}

	//Make workers and start working
	for i := 0; i < N; i++ {
		worker := makeWorker(ch, i)
		workers = append(workers, *worker)
		go worker.Start()
	}

	for i := 0; i < len(data); i++ {
		ch <- data[i]
	}

	time.Sleep(1 * time.Second)

	for _, v := range workers {
		v.Stop()
	}
}
