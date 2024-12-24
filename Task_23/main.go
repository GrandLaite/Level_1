/* Удалить i-ый элемент из слайса. */
package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}

	var i int
	fmt.Print("Введите индекс для удаления: ")
	fmt.Scan(&i)

	if i < 0 || i >= len(slice) {
		fmt.Println("Индекс вне диапозона.")
		return
	}

	slice = append(slice[:i], slice[i+1:]...)

	fmt.Println("Слайс после удаления:", slice)
}
