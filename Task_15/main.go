/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}
func main() {
  someFunc()
}

В justString = v[:100] мы создаём срез строки, который имеет первые 100 символов от исходной строки. При этом слайс создаёт ссылку на всю исходную строку
и поэтому в памяти хранится вся строка, что непоптимально с точки зрения расхода памяти. Для решения этой проблемы нам необходимо создать копию первых 100 символов, а не ссылаться на срез.
*/

package main

import "fmt"

func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFunc() string {
	v := createHugeString(1 << 10)
	substring := string([]byte(v[:100]))
	return substring
}

func main() {
	justString := someFunc()
	fmt.Println("Результат:", justString)
}
