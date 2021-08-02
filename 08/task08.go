package main

import "fmt"

func main() {
	var num int64
	num = 1
	fmt.Println("Number before setBit: ", num)
	num = setBit(num, 0, 0)
	fmt.Println("Number after setBit: ", num)
}

func setBit(num int64, pos uint, bit int) int64 {
	if bit == 1 {
		num |= (1 << pos)
		return num
	}
	if bit == 0 {
		var mask int64
		//xor to reverse the bits in mask
		mask = ^(1 << pos)
		num &= mask
		return num
	}
	return 0
}
