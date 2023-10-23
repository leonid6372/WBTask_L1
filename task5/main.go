// Task 5

package main

import (
	"fmt"
	"time"
)

// Send to data channel
func Send(data chan int, stop, done chan struct{}) {
	defer close(data) // Close data channel when Sender finishing

	// Send to data channel until stop channel will close
	i := 0
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-stop:
			done <- struct{}{} // Send correct finish to done channel
			return
		default:
			<-ticker.C
			data <- i
			i++
		}
	}
}

// Read from data channel
func Read(data chan int, stop, done chan struct{}) {
	// Read and print from data channel until stop channel will close
	for {
		select {
		case <-stop:
			done <- struct{}{} // Send correct finish to done channel
			return
		default:
			fmt.Println(<-data)
		}
	}
}

func main() {
	var secDuration int = 1
	fmt.Print("Enter program duration in seconds: ")
	fmt.Scan(&secDuration)

	data := make(chan int)
	stop, done := make(chan struct{}), make(chan struct{})

	// Wait until time will up, then close stop channel for Sender and Reader
	duration := time.After(time.Duration(secDuration) * time.Second)
	go func() {
		<-duration
		close(stop)
	}()

	// Start Sender and Reader concurrently
	go Send(data, stop, done)
	go Read(data, stop, done)

	// Wait for Sender and Reader are finished correct
	<-done
	<-done
}
