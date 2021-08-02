package main

import "fmt"

func update(p *int) {
	b := 2
	p = &b
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}

//Вывод данной программы: 1 1
//Все аргументы передаются в функцию по значению. При вызове update() создается локальная
//копия p, адрес которой изменяется внутри, при этом p в main() остается неизменной.
