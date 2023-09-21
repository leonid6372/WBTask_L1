// Task 1

package main

import (
	"fmt"
)

// Parent struct Human
type Human struct {
	Age int
}

// Parent struct method
func (h *Human) Greeting() {
	fmt.Printf("Hello! I'm human and %d years old", h.Age)
}

// Child struct
type Action struct {
	// Embed parent struct in child
	Human
}

func main() {
	// Create parent struct
	human := Human{Age: 23}

	// Create child struct
	action := Action{Human: human}

	// Access parent struct value from child struct
	if action.Age == 23 {
		// Access parent struct method from child struct
		action.Greeting()
	}
}
