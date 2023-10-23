// Task 4

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Reader from data channel
func Worker(n int, data chan int, stop, workerDone chan struct{}) {
	// Read and print from data channel until stop channel will close
	for {
		select {
		case <-stop:
			workerDone <- struct{}{} // Send correct finish to done channel
			return
		default:
			fmt.Printf("Worker %d: %d\n", n, <-data)
		}
	}
}

// Sender data to channel
func Send(data chan int, stop, senderDone chan struct{}) {
	defer close(data) // Close data channel when Sender finishing

	// Send to data channel until stop channel will close
	i := 0
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-stop:
			senderDone <- struct{}{} // Send correct finish to done channel
			return
		default:
			<-ticker.C
			data <- i
			i++
		}
	}
}

func main() {
	var N int
	fmt.Print("Enter workers amount to create: ")
	fmt.Scan(&N)

	data := make(chan int)
	stop := make(chan struct{})
	workerDone := make(chan struct{})
	senderDone := make(chan struct{})

	// Start N workers
	for i := 0; i < N; i++ {
		go Worker(i, data, stop, workerDone)
	}

	// Start data Sender
	go Send(data, stop, senderDone)

	// Program finishing when CTRL-C is pressed by user
	signalChan := make(chan os.Signal, 1)     // Channel for system signal
	signal.Notify(signalChan, syscall.SIGINT) // Catch SIGINT signal when it triggered by pressing user CTRL-C
	go func() {
		<-signalChan // Wait signal from signalChan
		close(stop)  // Close stop channel to stop sender and workers
	}()

	// Wait for Sender and N workers are finished correct
	<-senderDone
	for i := 0; i < N; i++ {
		<-workerDone
	}
}
