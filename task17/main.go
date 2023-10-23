// Task 17

package main

import (
	"fmt"
)

// Search recursive function
func BinarySearch(key int, arr []int, min, max int) int {
	// Check that key in array range
	if key > max || key < min {
		fmt.Printf("There is no element %v in array\n", key)
		return len(arr)
	}

	// Half array by mid index
	mid := (min + max) / 2
	if arr[mid] < key { // Check left side
		return BinarySearch(key, arr, mid, max)
	} else if arr[mid] > key { // Check right side
		return BinarySearch(key, arr, min, mid)
	} else if arr[mid] == key { // Check equality
		return mid
	} else {
		fmt.Printf("There is no element %v in array\n", key)
		return len(arr) // Return len(arr) if element does not exist
	}
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(BinarySearch(6, arr, 0, len(arr)))
}
