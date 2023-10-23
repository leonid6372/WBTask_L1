// Task 20

package main

import (
	"fmt"
)

// Split into words
func SplitIntoWords(str *string) []string {
	var words []string
	var word []rune

	// Add word to words if current cymbol is space and word is not empty
	for _, symbol := range *str {
		if symbol == ' ' {
			if len(word) != 0 {
				words = append(words, string(word))
				word = []rune{}
			}
		} else {
			word = append(word, symbol) // Add symbol to word if current symbol is not space
		}
	}

	// Add last word to words
	if len(word) != 0 {
		words = append(words, string(word))
	}

	return words
}

// TurnAround function
func TurnAround(words []string) {
	// Swap mirror strings
	N := (len(words)) - 1
	for i := 0; i <= N/2; i++ {
		words[i], words[N-i] = words[N-i], words[i]
	}
}

func main() {
	str := "   snow   dog sun  "

	words := SplitIntoWords(&str)
	TurnAround(words)

	fmt.Println(words)
}
