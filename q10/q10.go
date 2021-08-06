package main

import "fmt"

func main() {
	s1 := make([]int, 3)
	s2 := make([]int, 3, 4)
	s3 := []int{1, 2, 3}
	var s4 []int
	var s5 = []int{}
	var s6 = new([]int)

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
	fmt.Println(s4, len(s4), cap(s4))
	fmt.Println(s5, len(s5), cap(s5))
	fmt.Println(*s6, len(*s6), cap(*s6))

	m1 := make(map[int]int)
	m2 := make(map[int]int, 4)
	var m3 = map[int]int{}
	m4 := map[int]int{}
	fmt.Println(m1, len(m1))
	fmt.Println(m2, len(m2))
	fmt.Println(m3, len(m3))
	fmt.Println(m4, len(m4))

}
