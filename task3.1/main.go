// Task 3

package main

import (
	"fmt"
)

// Calculate square of number and send it to channel
func CalcSquare(number int, ch chan int) {
	square := number * number
	ch <- square
}

func main() {
	array := []int{2, 4, 6, 8, 10}
	sum := 0

	ch := make(chan int) // Make channel to synchronize Goroutines

	// Start Goroutine for every array element
	for _, value := range array {
		go CalcSquare(value, ch)
	}

	// Wait squares of all array elements from channel and calculate their sum
	for i := 0; i < len(array); i++ {
		sum += <-ch
	}

	// Print answer
	fmt.Println(sum)
}
