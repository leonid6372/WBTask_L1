// Task 21

package main

import (
	"fmt"
)

// Adaptee struct
type StringParser struct {
	words []string
}

// Adaptee struct function
func (sp *StringParser) PrintWords() {
	fmt.Println(sp.words)
}

// Print interface
type StringStruct interface {
	Print()
}

// Needed function to use with adaptee struct
func PrintString(ss StringStruct) {
	ss.Print()
}

// Adapter for adaptee struct StringParse
type Adapter struct {
	StringParser
}

// Adapter print function
func (a *Adapter) Print() {
	a.StringParser.PrintWords()
}

func main() {
	stringParser := StringParser{
		words: []string{"car", "holding", "key"}, // Create adaptee struct
	}
	adapter := Adapter{StringParser: stringParser} // Create adapter
	PrintString(&adapter)                          // Use PrintString() function with adaptee struct through adapter
}
