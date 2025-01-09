/* Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a и b, значение которых > 2^20. */
package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("1048577", 10)
	b.SetString("3097153", 10)

	sum := new(big.Int).Add(a, b)
	fmt.Printf("Сумма: %s\n", sum.String())

	sub := new(big.Int).Sub(a, b)
	fmt.Printf("Разность: %s\n", sub.String())

	mul := new(big.Int).Mul(a, b)
	fmt.Printf("Произведение: %s\n", mul.String())

	div := new(big.Rat).SetFrac(a, b)
	fmt.Printf("Частное: %s\n", div.FloatString(10))

}
