package main

import "fmt"

type Human struct {
	Name string
	Age  uint
}

type Action struct {
	Human
}

func (a *Action) printName() {
	fmt.Println(a.Name)
}

func main() {
	h := Human{
		Name: "Alex",
		Age:  25,
	}
	a := Action{h}
	a.printName()
}
