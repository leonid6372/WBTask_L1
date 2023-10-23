// Task 26

package main

import (
	"fmt"
	"sort"
	"strings"
)

// isUnique return true if every symbol in string is unique
func isUnique(str string) bool {
	// Convert all symbols to lowercase, cast string to slice of runes and sort slice
	lowCaseStr := strings.ToLower(str)
	runeStr := []rune(lowCaseStr)
	sort.Slice(runeStr, func(i, j int) bool {
		return runeStr[i] <= runeStr[j]
	})

	// Chek that every rune is unique
	for i := 0; i < len(runeStr)-1; i++ {
		if runeStr[i] == runeStr[i+1] {
			return false // If symbols repeats return false and quit function
		}
	}

	return true
}

func main() {
	strs := []string{"abcd", "abCdefAaf", "aabcd"}
	var res []bool
	for _, str := range strs {
		res = append(res, isUnique(str))
	}
	fmt.Println(res)
}
