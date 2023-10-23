// Task 7

package main

import (
	"fmt"
	"sync"
)

// Write to map while data channel is open
func Write(i int, testMap map[int]int, mtx *sync.Mutex, data chan int, done chan struct{}) {
	for temp := range data { // Avoid reading default value from closed channel
		mtx.Lock() // Lock to avoid writing data at the same time with other Writers
		testMap[temp] = temp
		mtx.Unlock() // Unlock after writing
	}
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	data := make(chan int)
	testMap := make(map[int]int)
	N, dataCount := 4, 15
	var mtx sync.Mutex // Create mutex to separate writing between Writers

	for i := 0; i < N; i++ {
		go Write(i, testMap, &mtx, data, done)
	}

	for i := 1; i < dataCount; i++ {
		data <- i
	}
	close(data)

	for i := 0; i < N; i++ {
		<-done
	}

	fmt.Println(testMap)
}
