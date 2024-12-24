/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	var duration int
	fmt.Print("Введите время работы программы: ")
	fmt.Scan(&duration)

	dataChannel := make(chan int)

	timer := time.NewTimer(time.Duration(duration) * time.Second)

	go func() {
		counter := 1
		for {
			select {
			case <-timer.C:
				close(dataChannel)
				return
			default:
				dataChannel <- counter
				counter++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for value := range dataChannel {
		fmt.Printf("Получено значение: %d\n", value)
	}

	fmt.Println("Программа завершена.")
}
