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
