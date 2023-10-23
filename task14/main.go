// Task 14

package main

import (
	"fmt"
)

func main() {
	var test interface{}
	fmt.Scan(&test)

	// Define data type with switch construction
	switch testType := test.(type) {
	case int:
		fmt.Printf("Data type: %T. Value: %v.\n", testType, test)
	case string:
		fmt.Printf("Data type: %T. Value: %v.\n", testType, test)
	case bool:
		fmt.Printf("Data type: %T. Value: %v.\n", testType, test)
	case chan interface{}:
		fmt.Printf("Data type: %T. Value: %v.\n", testType, test)
	default:
		fmt.Printf("I don't know about type %T!\n", testType)
	}
}
