/* Реализовать конкурентную запись данных в map. */

package main

import (
	"fmt"
	"sync"
)

func main() {
	dataMap := make(map[int]string)
	var mu sync.Mutex
	var wg sync.WaitGroup

	numWorkers := 10

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			mu.Lock()
			dataMap[workerID] = fmt.Sprintf("Данные из воркера №%d", workerID)
			mu.Unlock()
			fmt.Printf("Горутина %d записала данные в map.\n", workerID)
		}(i)
	}

	wg.Wait()

	fmt.Println("Результат записи в map:")
	for key, value := range dataMap {
		fmt.Printf("Ключ: %d, Значение: %s\n", key, value)
	}
}
