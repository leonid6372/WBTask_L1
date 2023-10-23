// Task 18

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Concurrent increment function
func Increment(i *int32, wg *sync.WaitGroup) {
	defer wg.Done()       // WaitGroup done
	atomic.AddInt32(i, 1) // Use atomic.Add... to increase counter concurrently
}

func main() {
	var res int32 = 0

	wg := new(sync.WaitGroup)
	wg.Add(100)

	// Start concurrent counter
	for i := 0; i < 100; i++ {
		go Increment(&res, wg)
	}

	wg.Wait()

	fmt.Println(res)
}
