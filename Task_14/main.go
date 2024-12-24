/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/
package main

import "fmt"

func main() {
	var1 := 51692
	var2 := "Argooooon!"
	var3 := false
	var4 := make(chan int)

	checkType(var1)
	checkType(var2)
	checkType(var3)
	checkType(var4)
}

func checkType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("Переменная %v имеет тип int\n", v)
	case string:
		fmt.Printf("Переменная \"%v\" имеет тип string\n", v)
	case bool:
		fmt.Printf("Переменная %v имеет тип bool\n", v)
	case chan int:
		fmt.Println("Переменная является каналом типа int")
	default:
		fmt.Printf("Тип переменной неизвестен: %T\n", v)
	}
}
