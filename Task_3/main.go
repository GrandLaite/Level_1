/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	sum := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			square := n * n
			mu.Lock()
			sum += square
			mu.Unlock()
		}(num)
	}

	wg.Wait()

	fmt.Printf("Сумма квадратов чисел: %d\n", sum)
}
