package main

import "fmt"

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}

//[b b a] [a a]
//Внутри функции создается локальная копия slice'a. При append'e количество элементов превышает cap слайса.
//Создается новый внутренний массив для элементов. Слайс снаружи указывает внутри на тот же массив, что и раньше.
//Из-за этого разный вывод.
