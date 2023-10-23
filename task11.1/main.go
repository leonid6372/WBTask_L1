// Task 11

package main

import (
	"fmt"
)

// Find intersection
func GetIntersection(set1, set2 map[int]struct{}) map[int]struct{} {
	intersection := make(map[int]struct{})

	for key1, _ := range set1 {
		for key2, _ := range set2 {
			if key1 == key2 {
				intersection[key1] = struct{}{}
			}
		}
	}

	return intersection
}

func main() {
	set1 := map[int]struct{}{
		10: struct{}{},
		2:  struct{}{},
		6:  struct{}{},
		4:  struct{}{},
		11: struct{}{},
		12: struct{}{},
		14: struct{}{},
	}
	set2 := map[int]struct{}{
		5:  struct{}{},
		1:  struct{}{},
		8:  struct{}{},
		2:  struct{}{},
		9:  struct{}{},
		3:  struct{}{},
		4:  struct{}{},
		10: struct{}{},
		0:  struct{}{},
		6:  struct{}{},
		7:  struct{}{},
	}

	for key, _ := range GetIntersection(set1, set2) {
		fmt.Printf("%d ", key)
	}
}
