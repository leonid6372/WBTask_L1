// Task 23

package main

import (
	"fmt"
	"reflect"
)

// Erase from slice of any data type
func Erase(s interface{}, i int) {
	// Get value of s (pointer) and get value that s point to
	// Check that input slice length not less than element index to erase
	if v := reflect.Indirect(reflect.ValueOf(s)); v.Len() > i {
		reflect.AppendSlice(v.Slice(0, i), v.Slice(i+1, v.Len())) // Reset slice without erased element
		v.SetLen(v.Len() - 1)                                     // Reset slice length
	}
}

func main() {
	//words := []string{"car", "home", "windows", "port"}
	nums := []int{1, 2, 3, 4}
	fmt.Println(nums)
	Erase(&nums, 0)
	fmt.Println(nums)
}
