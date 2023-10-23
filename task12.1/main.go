// Task 12

package main

import (
	"fmt"
)

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	// Make map with empry struct value as set
	set := make(map[string]struct{})

	// Fill map from slice
	for _, str := range slice {
		set[str] = struct{}{}
	}

	// Print result
	for str := range set {
		fmt.Printf("%s ", str)
	}
}
