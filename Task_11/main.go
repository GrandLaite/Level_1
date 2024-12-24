/*
Реализовать пересечение двух неупорядоченных множеств.
*/
package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{3, 4, 5, 6, 7}

	numsMap := make(map[int]bool)
	for _, value := range nums1 {
		numsMap[value] = true
	}

	intersection := []int{}
	for _, value := range nums2 {
		if numsMap[value] {
			intersection = append(intersection, value)
		}
	}

	fmt.Println("Пересечение множеств:", intersection)
}
