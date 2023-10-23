// Task 10

package main

import (
	"fmt"
	"sync"
)

// Sort by decimicial of temperature
func Sort(data chan float32, sortedTemperatures map[int][]float32, mtx *sync.Mutex, done chan struct{}) {
	for temperature := range data { // Read from data channel while it is open
		decTemperature := int(temperature) / 10 * 10
		mtx.Lock()
		sortedTemperatures[decTemperature] = append(sortedTemperatures[decTemperature], temperature)
		mtx.Unlock()
	}
	done <- struct{}{} // Correct finish
}

func main() {
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	sortedTemperatures := make(map[int][]float32)
	N := 4
	data := make(chan float32)
	done := make(chan struct{})
	var mtx sync.Mutex

	// Start concurrent sorting
	for i := 0; i < N; i++ {
		go Sort(data, sortedTemperatures, &mtx, done)
	}

	// Send temperatures to Sort through data channel
	for _, temperature := range temperatures {
		data <- temperature
	}
	close(data)

	// Wait until all Sort will finish
	for i := 0; i < N; i++ {
		<-done
	}

	fmt.Println(sortedTemperatures)
}
