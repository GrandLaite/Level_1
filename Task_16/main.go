/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]

	left := []int{}
	right := []int{}
	middle := []int{}

	for _, value := range arr {
		if value < pivot {
			left = append(left, value)
		} else if value > pivot {
			right = append(right, value)
		} else {
			middle = append(middle, value)
		}
	}

	return append(append(quickSort(left), middle...), quickSort(right)...)
}

func main() {
	arr := []int{10, -1, 2, 5, 0, 6, 4, -5, 3}

	fmt.Println("Исходный массив:", arr)

	sorted := quickSort(arr)

	fmt.Println("Отсортированный массив:", sorted)
}
