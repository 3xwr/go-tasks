package main

import "fmt"

func main() {
	arr := make([]int, 100)
	fmt.Println(cap(arr))
}
