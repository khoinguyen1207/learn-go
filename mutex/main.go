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
		wg.Go(func() {
			mu.Lock()
			count++
			mu.Unlock()
		})
	}

	wg.Wait()

	fmt.Println("Count: ", count)
}
