// Task 22

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var strA, strB, operation string

	// Read a and b in string
	fmt.Print("Enter number A: ")
	fmt.Scan(&strA)
	fmt.Print("Enter number B: ")
	fmt.Scan(&strB)

	// Convert a and b from string to big.Int
	bigA := new(big.Int)
	bigA.SetString(strA, 10)
	bigB := new(big.Int)
	bigB.SetString(strB, 10)
	bigRes := new(big.Int)

	// Read operation and calculate answer
	for {
		fmt.Print("Enter operation sign (+ - * /): ")
		fmt.Scan(&operation)

		switch operation {
		case "+":
			bigRes.Add(bigA, bigB)
			fmt.Printf("Answer: " + bigRes.String())
			return
		case "-":
			bigRes.Sub(bigA, bigB)
			fmt.Printf("Answer: " + bigRes.String())
			return
		case "*":
			bigRes.Mul(bigA, bigB)
			fmt.Printf("Answer: " + bigRes.String())
			return
		case "/":
			bigRes.Div(bigA, bigB)
			fmt.Printf("Answer: " + bigRes.String())
			return
		default:
			fmt.Println("Entered sign is wrong. Try again.")
		}
	}
}
