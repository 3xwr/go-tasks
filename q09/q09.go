package main

import "fmt"

func main() {
	//make используется для только создания слайса, мапы или канала, а с помощью new можно так же
	//создать переменные других типов
	//new просто выделяет память, в то время как make выделяет и иницилизирует память.
	//new возвращает указатель а make - тип
	a := new(int)
	//b := make(int) - error
	fmt.Println(a, *a)
	b := make(map[int]int)
	c := new(map[int]int)
	d := make(chan int, 5)
	fmt.Println(b)
	fmt.Println(*c)
	fmt.Println(d)
}
