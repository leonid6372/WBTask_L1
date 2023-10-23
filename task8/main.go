// Task 8

package main

import (
	"fmt"
	"log"
	"strconv"
)

// Change selected bit function
func ChangeBitInt64(number int64, bit, value int) int64 {
	// Check exception with int64 min value
	if (number == 0 || number == -9223372036854775808) && bit == 1 && value == 1 {
		return -9223372036854775808
	}

	// Convert int64 to []rune with binary number		Example: 5 -> ['1' '0' '1']
	binNumber := []rune(fmt.Sprintf("%b", uint64(number)))

	// If number >= 0 add runes '0' in front of significant digits to get summary 64 runes
	// If number < 0 runes slice already had 64 runes
	if number >= 0 {
		var temp []rune
		for i := 0; i < 64-len(binNumber); i++ {
			temp = append(temp, '0')
		}
		temp = append(temp, binNumber...)
		binNumber = temp
	}

	// Make var rune for bit value
	var newValue rune
	if value == 1 {
		newValue = '1'
	} else {
		newValue = '0'
	}

	// Change selected bit in runes slice
	binNumber[bit-1] = newValue

	// Check new number < 0 (1st bit == 1)
	isMinus := false
	if binNumber[0] == '1' {
		isMinus = true

		// For negative numbers do next steps to convert it to decimicial

		// Pass 1st bit and find first 1 to subtract 1 from binary number
		for i := 63; i > 0; i-- {
			if binNumber[i] == '1' {
				binNumber[i] = '0'            // Change 1 by 0
				for j := i + 1; j < 64; j++ { // Change all next bits by 0
					binNumber[j] = '1'
				}
				break
			}
		}
		// Negation all bits except sign bit
		for i := 1; i < 64; i++ {
			if binNumber[i] == '1' {
				binNumber[i] = '0'
			} else {
				binNumber[i] = '1'
			}
		}
	}

	// Here new binary number with 1st sign bit and runes '0' in front of significant digits
	for i := 1; i < 64; i++ {
		if binNumber[i] == '1' {
			binNumber = binNumber[i:] // Set only significant digits
			break
		}
	}

	// Convert binary number in []rune to string to decimicial int64
	number, err := strconv.ParseInt(string(binNumber), 2, 64)
	if err != nil {
		log.Println(err)
	}

	// Multiply by -1 to get minus if isMinus == true
	if isMinus {
		return (number * -1)
	} else {
		return number
	}
}

func main() {
	var number int64 = 0
	var bit, value int = 56, 1
	newNumber := ChangeBitInt64(number, bit, value)
	fmt.Println(newNumber)
}
