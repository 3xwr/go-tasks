package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

// func main() { --- исправленная версия программы
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(wg *sync.WaitGroup, i int) {
// 			fmt.Println(i)
// 			wg.Done()
// 		}(&wg, i)
// 	}
// 	wg.Wait()
// 	fmt.Println("exit")
// }

//Произойдет паника, так как создается копия вейтгруппы. Для того, чтобы программа работала корректно
//нужно передать вейтгруппу по ссылке.
