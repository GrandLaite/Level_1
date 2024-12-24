/* Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout. */

package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	firstChannel := make(chan int)
	secondChannel := make(chan int)

	go func() {
		for _, num := range numbers {
			firstChannel <- num
		}
		close(firstChannel)
	}()

	go func() {
		for num := range firstChannel {
			secondChannel <- num * 2
		}
		close(secondChannel)
	}()

	for result := range secondChannel {
		fmt.Printf("Результат: %d\n", result)
	}
}
