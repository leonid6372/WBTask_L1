// Task 11

package main

import (
	"fmt"
)

// Find intersection
func GetIntersection(array1, array2 []int) []int {
	var intersection []int

	for _, value1 := range array1 {
		for _, value2 := range array2 {
			if value1 == value2 {
				intersection = append(intersection, value1)
			}
		}
	}

	return intersection
}

func main() {
	array1 := []int{10, 2, 6, 4, 11, 12, 14}
	array2 := []int{5, 1, 8, 2, 9, 3, 4, 10, 0, 6, 7}
	fmt.Println(GetIntersection(array1, array2))
}
