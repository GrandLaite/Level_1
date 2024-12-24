/*
Реализовать все возможные способы остановки выполнения горутины.
*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// через канал
	wg.Add(1)
	stopChan := make(chan bool)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopChan:
				fmt.Println("Горутина 1 завершена через канал.")
				return
			default:
				fmt.Println("Горутина 1 выполняется.")
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	stopChan <- true

	// через контекст
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина 2 завершена через контекст.")
				return
			default:
				fmt.Println("Горутина 2 выполняется.")
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	cancel()

	// через флаг
	wg.Add(1)
	stopFlag := false
	go func() {
		defer wg.Done()
		for {
			if stopFlag {
				fmt.Println("Горутина 3 завершена через флаг.")
				return
			}
			fmt.Println("Горутина 3 выполняется.")
			time.Sleep(1000 * time.Millisecond)
		}
	}()
	time.Sleep(3 * time.Second)
	stopFlag = true

	// через таймаут
	wg.Add(1)
	go func() {
		defer wg.Done()
		timer := time.NewTimer(2 * time.Second)
		for {
			select {
			case <-timer.C:
				fmt.Println("Горутина 4 завершена через таймаут.")
				return
			default:
				fmt.Println("Горутина 4 выполняется.")
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}
