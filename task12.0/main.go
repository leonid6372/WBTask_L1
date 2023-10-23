// Task 12

package main

import (
	"fmt"
	"sort"
)

// Make struct with set properties
type StringSet struct {
	value []string
}

// Constructor for struct
func NewStringSet(strings []string) StringSet {
	s := new(StringSet)

	// Sort
	sort.Strings(strings)

	// Unique
	s.value = append(s.value, (strings)[0])
	for i := 0; i < len(strings)-1; i++ {
		if strings[i] != strings[i+1] {
			s.value = append(s.value, (strings)[i+1])
		}
	}

	return *s
}

// Insert method returning inserted element index
func (s *StringSet) Insert(outterStr string) int {
	for i, innerStr := range s.value {
		if outterStr == innerStr {
			return len(s.value)
		}
		if outterStr < innerStr {
			s.value = append(s.value, "")
			copy(s.value[i+1:], s.value[i:])
			s.value[i] = outterStr
			return i
		}
	}
	s.value = append(s.value, outterStr)
	return len(s.value)
}

// Erace method returning element index before erased element
func (s *StringSet) Erase(outterStr string) int {
	for i, innerStr := range s.value {
		if outterStr == innerStr {
			copy(s.value[i:], s.value[i+1:])
			//s.value[len(s.value)-1] = ""
			s.value = s.value[:len(s.value)-1]
			return i
		}
		if outterStr < innerStr {
			return i
		}
	}
	return len(s.value)
}

// Find method returning found element index
func (s *StringSet) Find(outterStr string) int {
	for i, innerStr := range s.value {
		if outterStr == innerStr {
			return i
		}
		if outterStr < innerStr {
			break
		}
	}
	return len(s.value)
}

// Iterate method using in "for ... range" constriction
func (s *StringSet) Iterate() []string {
	return s.value
}

func main() {
	// Make set struct by constructor
	testStringSet := NewStringSet([]string{"cat", "cat", "dog", "cat", "tree"})

	// Test insert method
	testStringSet.Insert("apple")
	testStringSet.Insert("zone")
	testStringSet.Insert("zip")

	// Test iterate method
	for i, str := range testStringSet.Iterate() {
		fmt.Println(i, str)
	}

	fmt.Println("-------------------------")

	// Test erase method
	it := testStringSet.Erase("dog")
	fmt.Printf("Position %d was erased from set\n", it)

	for i, str := range testStringSet.Iterate() {
		fmt.Println(i, str)
	}

	// Test find method
	fmt.Println(testStringSet.Find("tree"), testStringSet.Find("car"))
}
