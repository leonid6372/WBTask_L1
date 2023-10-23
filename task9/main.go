// Task 9

package main

import (
	"fmt"
)

// 1st pipeline part
func Read(rawData chan int, buffer chan int) {
	defer close(buffer)

	// Read data and send to next pipeline part while rawData is open
	for number := range rawData {
		buffer <- number
	}
}

// 2nd pipeline part
func Print(buffer chan int, pipelineDone chan struct{}) {

	// Muliply by 2 and print result in stdout while buffer is open
	for newNumber := range buffer {
		fmt.Printf("%d ", newNumber)
	}
	pipelineDone <- struct{}{}
}

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rawData := make(chan int)
	buffer := make(chan int)
	pipelineDone := make(chan struct{})

	// 1st pipeline part
	go Read(rawData, buffer)

	// 2nd pipeline part
	go Print(buffer, pipelineDone)

	// Send data to pipeline
	for _, value := range array {
		rawData <- value
	}
	close(rawData) // All data sent. Close data channel

	// Wait pipeline finishing
	<-pipelineDone
}
