package main

import (
	"fmt"
)

// Main logic at line 22 (i.e if len(stack) == 0 || stack[len(stack)-1] != '{'  ). 
// Current character is '}' which is an opening bracket => the immediate previous element
// must be '{' for this to be a valid string because we are not pushing any other character in the string
// and also the brackets open and close in LIFO order. So returning false if either stack is empty OR last element is not equal to '{'
// Likewise for other brackets also at line 27,32

func checkValidBrackets(s string) bool {

	stack := []rune{} // This is a slice which will be used as stack

	for _, ch := range s {
		switch ch {
		case '{', '(', '[':
			stack = append(stack, ch) // if character is either of the opening bracket - we pust it on to the slice(stack)
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1] // this operation of removing last element from slice acts as pop operation here
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0 // i.e if len(stack) == 0, return true else return false
}
func main() {
	listStrings := []string{
		"()",
		"([]){}",
		"(]",
		"([)",
		"{[()]}",
		"(((",
		"",
	}
	for _, str := range listStrings {
		//fmt.Println("String is", str)
		out := checkValidBrackets(str)
		fmt.Printf("Output for validity of string %s is %v\n", str, out)
	}

}
