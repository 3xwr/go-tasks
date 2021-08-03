package main

import "fmt"

func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}

//Программа выведет 0, потому что внутри if true {} создается локальная переменная n, не имеющая ничего общего
//с n в теле функции main.
