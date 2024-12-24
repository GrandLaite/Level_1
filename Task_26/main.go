/* Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.


Например:

abcd — true

abCdefAaf — false

aabcd — false
*/

package main

import (
	"fmt"
	"strings"
)

func isUnique(input string) bool {
	input = strings.ToLower(input)

	charMap := make(map[rune]bool)

	for _, char := range input {
		if charMap[char] {
			return false
		}
		charMap[char] = true
	}

	return true
}

func main() {
	stringsToTest := []string{"abcd", "abCdefAaf", "aabcd"}
	for _, str := range stringsToTest {
		fmt.Printf("Строка \"%s\" имеет уникальные символы: %t\n", str, isUnique(str))
	}
}
