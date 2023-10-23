// Task 13

package main

import (
	"fmt"
)

func main() {
	a, b := 0, 1
	fmt.Printf("a = %d; b = %d\n", a, b)

	a, b = b, a
	fmt.Printf("a = %d; b = %d\n", a, b)
}
