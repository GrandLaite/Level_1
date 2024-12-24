/* Реализовать собственную функцию sleep. */
package main

import (
	"fmt"
	"time"
)

func Sleep(seconds int) {
	select {
	case <-time.After(time.Duration(seconds) * time.Second):
	}
}

func main() {
	fmt.Println("Начало программы.")
	fmt.Println("Ожидаем...")

	Sleep(10)

	fmt.Println("Время вышло.")
}
