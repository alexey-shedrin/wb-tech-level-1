package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Int).SetString("10000000000000000000", 10)
	b, _ := new(big.Int).SetString("20000000000000000000", 10)

	mulResult := new(big.Int)
	mulResult.Mul(a, b)
	fmt.Printf("Умножение (%s * %s) = %s\n", a, b, mulResult)

	divResult := new(big.Int)
	divResult.Div(b, a)
	fmt.Printf("Деление (%s / %s) = %s\n", b, a, divResult)

	sumResult := new(big.Int)
	sumResult.Add(a, b)
	fmt.Printf("Сложение (%s + %s) = %s\n", a, b, sumResult)

	diffResult := new(big.Int)
	diffResult.Sub(b, a)
	fmt.Printf("Вычитание (%s - %s) = %s\n", b, a, diffResult)

}
