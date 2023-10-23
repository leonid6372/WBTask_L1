// Task 3

package main

import (
	"fmt"
	"sync"
)

// Function for concurrently calculating squares of numbers and their sum
// It's safe function to concurrently execute
// Every Goroutine will wait reading from channel until other send updated sum there
func UpdateSum(ch chan int, number int, wg *sync.WaitGroup) {
	defer wg.Done()           // This one of WaitGroup done
	square := number * number // Calculate square of number
	sum := <-ch + square      // Read current sum from channel
	ch <- sum                 // Send updated sum to channel
}

func main() {
	array := []int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup) // Create WaitGroup for waiting all Goroutines finishing in main()

	ch := make(chan int, 1) // Create channel with buffer by 1 int for access to sum value between Goroutines
	ch <- 0                 // Send start value in channel

	// Update sum for every number of array concurrently
	for _, value := range array {
		wg.Add(1)                   // Increase WaitGroup by 1
		go UpdateSum(ch, value, wg) // Start Goroutine
	}

	wg.Wait() // Wait until whole WaitGroup will be done
	close(ch) // Close channel

	fmt.Println(<-ch) // Print sum from channel
}
