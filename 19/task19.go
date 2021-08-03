package main

import "fmt"

// var justString string

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
		v += "A"
	}
	return v
}

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

func someFuncFixed(justString *string) {
	v := createHugeString(1 << 10)
	*justString = v[:100]
}

func main() {
	//someFunc()
	var justString string
	someFuncFixed(&justString)
	fmt.Println(justString)
}

//1. Глобальная переменная justString может изменяться везде внутри пакета, что может привести к
//неопределенному поведению. Например, можно далее объявить локальную переменную с именем justString
//другого типа, в связи с чем может произойти ошибка при попытке работы с justString как со строкой.
