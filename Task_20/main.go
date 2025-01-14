/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseWords(input string) string {
	words := strings.Fields(input)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	reversed := reverseWords(input)
	fmt.Printf("Перевёрнутая строка: %s\n", reversed)
}
