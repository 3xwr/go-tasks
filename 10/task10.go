package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	setMap := make(map[int][]float64)
	addNumsToSet(arr, setMap)
	fmt.Println(setMap)
}

func addNumsToSet(arr []float64, setMap map[int][]float64) {
	for _, v := range arr {
		num := int(math.Floor(v))
		//divide by ten until we get the first digit of number
		for num/10 != 0 {
			num /= 10
		}
		//multiply by 10 to get the 10s groups
		num *= 10
		setMap[num] = append(setMap[num], v)
	}
}
