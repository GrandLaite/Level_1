/* Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0. */

package main

import (
	"fmt"
)

func main() {
	var number int64
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	var bitIndex int
	fmt.Print("Введите индекс бита (0-63): ")
	fmt.Scan(&bitIndex)

	var operation int
	fmt.Print("Введите 1 для установки бита в 1 или 0 для установки в 0: ")
	fmt.Scan(&operation)

	if bitIndex < 0 || bitIndex > 63 {
		fmt.Println("Ошибка: индекс бита должен быть в пределах от 0 до 63.")
		return
	}

	switch operation {
	case 1:
		number = number | (1 << bitIndex)
		fmt.Printf("Результат после установки бита %d в 1: %d\n", bitIndex, number)
	case 0:
		number = number & ^(1 << bitIndex)
		fmt.Printf("Результат после установки бита %d в 0: %d\n", bitIndex, number)
	default:
		fmt.Println("Операция должна быть либо 1, либо 0.")
	}
}
