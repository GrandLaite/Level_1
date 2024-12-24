/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})
	for _, value := range sequence {
		set[value] = struct{}{}
	}

	uniqueValues := []string{}
	for key := range set {
		uniqueValues = append(uniqueValues, key)
	}

	fmt.Println(uniqueValues)
}
