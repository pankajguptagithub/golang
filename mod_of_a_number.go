// Code to find out modulus or absolute value of a given number 

package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}

}
func main() {
	n := -4659
	mod := abs(n)
	fmt.Printf("Back to main. mod of %d is %d\n", n, mod)
}
