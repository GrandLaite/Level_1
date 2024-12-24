/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	dataChannel := make(chan int)

	var workerCount int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&workerCount)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(dataChannel)
				return
			default:
				dataChannel <- rand.Intn(100)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	var wg sync.WaitGroup
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(ctx, &wg, dataChannel, i)
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	<-signalChannel
	fmt.Println("\nЗавершается работа...")

	cancel()

	wg.Wait()
	fmt.Println("Все воркеры завершили работу.")
}

func worker(ctx context.Context, wg *sync.WaitGroup, dataChannel <-chan int, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d завершает работу.\n", id)
			return
		case data, ok := <-dataChannel:
			if !ok {
				fmt.Printf("Воркер %d: канал закрыт, завершение.\n", id)
				return
			}
			fmt.Printf("Воркер %d получил данные: %d\n", id, data)
		}
	}
}
