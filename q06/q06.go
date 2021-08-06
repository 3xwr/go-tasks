package main

import "fmt"

func main() {
	var a []int
	b := []int{}

	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
	//Оба способа дают эквивалентную структуру, но второй способ вызова позволяет сразу добавить элементы
	//например:
	c := []int{1, 2, 3}
	fmt.Println(c, len(c), cap(c))
	//var d []int{1,2,3} - ошибка
}
