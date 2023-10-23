// Task 2

package main

import (
	"fmt"
	"sync"
)

// Function for concurrently printing squares
func PrintSquare(number int, wg *sync.WaitGroup) {
	defer wg.Done() // This one of WaitGroup done
	fmt.Printf("%d ", number*number)
}

func main() {
	array := []int{2, 4, 6, 8, 10}

	wg := new(sync.WaitGroup) // Create WaitGroup for waiting all Goroutines finishing in main()

	// Print square of every number of array concurrently
	for _, value := range array {
		wg.Add(1)                 // Increase WaitGroup by 1
		go PrintSquare(value, wg) // Start Goroutine
	}

	wg.Wait() // Wait until whole WaitGroup will be done
}
