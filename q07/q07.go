package main

import "fmt"

func showType(i interface{}) {
	fmt.Printf("%v is of type %T\n", i, i)
}

func main() {
	showType(1)
	showType(true)
}
