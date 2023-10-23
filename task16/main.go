// Task 16

package main

import (
	"fmt"
)

// Sort recursive function
func QuickSort(arr []int) {
	// Array of one element is already sorted -> quit
	if len(arr) == 1 {
		return
	}

	// Separate all elements less than pivot left from wall and all elements higher than pivot right from wall
	pivot := len(arr) - 1
	wall := 0
	for cur := wall; cur < pivot; cur++ {
		if arr[cur] <= arr[pivot] {
			arr[wall], arr[cur] = arr[cur], arr[wall]
			wall++
		}
	}
	arr[wall], arr[pivot] = arr[pivot], arr[wall]

	// Recursive sort for both parts sides of wall
	QuickSort(arr[:wall])
	QuickSort(arr[wall:])
}

func main() {
	arr := []int{2, 7, 3, 0, 9, 5, 4, 1, 8, 6}
	QuickSort(arr)
	fmt.Println(arr)
}
