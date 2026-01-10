package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			count++
			mu.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Count: ", count)
}
