/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
*/
package main

import (
	"fmt"
)

func reverseString(input string) string {
	runes := []rune(input)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	return string(runes)
}

func main() {
	var input string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)

	reversed := reverseString(input)

	fmt.Printf("Перевёрнутая строка: %s\n", reversed)
}
