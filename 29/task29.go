package main

import (
	"fmt"
	"math/big"
)

func Pow(a *big.Float, e uint64) *big.Float {
	result := big.NewFloat(0.0).Copy(a)
	for i := uint64(0); i < e-1; i++ {
		result = new(big.Float).Mul(result, a)
	}
	return result
}

func AddAndPrint(a *big.Float, b *big.Float) {
	addResult := new(big.Float).Add(a, b)
	addResult.SetPrec(256)
	fmt.Println(addResult)
}

func SubAndPrint(a *big.Float, b *big.Float) {
	subResult := new(big.Float).Sub(a, b)
	fmt.Println(subResult)
	subResult.SetPrec(256)
}

func MulAndPrint(a *big.Float, b *big.Float) {
	mulResult := new(big.Float).Mul(a, b)
	fmt.Println(mulResult)
	mulResult.SetPrec(256)
}

func DivAndPrint(a *big.Float, b *big.Float) {
	divResult := new(big.Float).Quo(a, b)
	fmt.Println(divResult)
	divResult.SetPrec(256)
}

func main() {
	//2^50
	bigNum := Pow(big.NewFloat(2), 21)
	num1 := bigNum
	num2 := bigNum
	AddAndPrint(num1, num2)
	SubAndPrint(num1, num2)
	MulAndPrint(num1, num2)
	DivAndPrint(num1, num2)
}
