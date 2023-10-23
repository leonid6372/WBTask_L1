// Task 6

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	stop, done := make(chan struct{}), make(chan struct{})

	// Way to stop goroutine №1. Send stop signal to channel
	go func() {
		fmt.Println("Goroutine0 started")
		for {
			select {
			case <-stop:
				done <- struct{}{}
				return
			default:
				time.Sleep(900 * time.Millisecond)
				fmt.Println("Goroutine0 is working")
			}
		}
	}()
	time.Sleep(2 * time.Second)
	stop <- struct{}{}
	<-done
	fmt.Println("Goroutine0 finished by sending signal to stop channel")

	// Way to stop goroutine №2. Close channel
	go func() {
		fmt.Println("Goroutine1 started")
		for {
			select {
			case <-stop:
				done <- struct{}{}
				return
			default:
				time.Sleep(900 * time.Millisecond)
				fmt.Println("Goroutine1 is working")
			}
		}
	}()
	time.Sleep(2 * time.Second)
	close(stop)
	<-done
	fmt.Println("Goroutine1 finished by closing stop channel")

	// Way to stop goroutine №3. Use context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go func() {
		fmt.Println("Goroutine2 started")
		for {
			select {
			case <-ctx.Done():
				done <- struct{}{}
				return
			default:
				time.Sleep(900 * time.Millisecond)
				fmt.Println("Goroutine2 is working")
			}
		}
	}()
	<-done
	fmt.Println("Goroutine2 finished by context Done()")
}
