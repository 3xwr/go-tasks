package main

import "fmt"

func main() {
	var i int
	var s string
	var b bool
	var ch chan int
	getType1(i)
	getType1(s)
	getType1(b)
	getType1(ch)
	fmt.Println("-------")
	getType2(i)
	getType2(s)
	getType2(b)
	getType2(ch)
}

func getType1(i interface{}) {
	fmt.Printf("%v is of type %T\n", i, i)
}

func getType2(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v is of type int\n", v)
	case string:
		fmt.Printf("%v is of type string\n", v)
	case bool:
		fmt.Printf("%v is of type bool\n", v)
	case chan int:
		fmt.Printf("%v is of type chan int\n", v)
	default:
		fmt.Println("Unknown type")
	}
}
