// Task 19

package main

import (
	"fmt"
)

// TurnAround function
func TurnAround(str *string) {
	r_str := []rune(*str)

	// Swap mirror runes
	N := (len(r_str)) - 1
	for i := 0; i <= N/2; i++ {
		r_str[i], r_str[N-i] = r_str[N-i], r_str[i]
	}

	*str = string(r_str)
}

func main() {
	str := "главрыба"
	TurnAround(&str)
	fmt.Println(str)
}
