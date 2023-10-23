// Task 25

package main

import (
	"fmt"
	"time"
)

// Sleep through time.After
func MySleep1(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second) // Wait for signal from channel time.After()
}

// Sleep through time.Ticker
func MySleep2(seconds int) {
	ticker := time.NewTicker(1 * time.Second)
	for i := 0; i < seconds; i++ {
		<-ticker.C
	}
}

func main() {
	duration := 3
	fmt.Println(time.Now())
	MySleep1(duration)
	fmt.Println(time.Now())
	MySleep2(duration)
	fmt.Println(time.Now())
}
