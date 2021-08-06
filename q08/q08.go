package main

import "fmt"

func main() {
	testMap := make(map[int]int)
	for i := 0; i < 10; i++ {
		testMap[i] = i
	}
	for i := 0; i < 5; i++ {
		for _, v := range testMap {
			fmt.Println(v)
		}
		fmt.Println("_______________")
	}
	//Порядок неопределенный.
}
