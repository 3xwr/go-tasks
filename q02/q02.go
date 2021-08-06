package main

import "fmt"

type Animal interface {
	Speak()
}

type Cow struct {
}

type Dog struct {
}

func (c *Cow) Speak() {
	fmt.Println("Moo!")
}

func (d *Dog) Speak() {
	fmt.Println("Bark!")
}

func main() {
	//В данном случае типы Cow и Dog реализуют интерфейс Animal.
	var c Cow
	var d Dog
	c.Speak()
	d.Speak()
}
