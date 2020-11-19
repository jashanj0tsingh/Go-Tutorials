package main

import "fmt"

// You are given a string that contains left and right parenthesis characters.
// Write code to determine whether the parentheses are correctly nested.
// For example, the strings "(())" and "()()" are correctly nested but "(()()" and ")(" are not.

func checkParenthesisNesting(str string) bool {

	var counter = 0
	// convert string to a slice of runes
	cs := []rune(str)

	for _, c := range cs {
		if c == '(' {
			counter++
		} else if c == ')' {
			counter--
			if counter < 0 { return false }
		}
	}

	return false
}

func main() {
	balanced := checkParenthesisNesting(")(")
	if balanced {
		fmt.Println("the parenthesis are balanced")
	} else {
		fmt.Println("the parenthesis are not balanced")
	}
}