// Code to reverse a given number 

package main

import "fmt"

func reverse(n int) int {	
	rev := 0
	for i := n; i > 0; i = i / 10 {	
		if i < 10 {			
			rev += i
		} else {
			r := i % 10
			fmt.Println("rev is ", rev)
			fmt.Println("r is ", r)
			rev = 10 * (r + rev)
		}
		fmt.Println("Reverse is ", rev)
		fmt.Println("===================")
	}
	return rev
}
func main() {
	n := 4659
	rev := reverse(n)
	fmt.Printf("Back to main. Reverse of %d is %d\n", n, rev)
}
