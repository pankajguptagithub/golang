package main

import (
	"fmt"
)

/*
A simple program to demonstrate how receivers work in Golang
Value receivers like in this example work on both value and pointer variables:
x := customInt(4)
x.customPrint()   // works
(&x).customPrint() // also works
*/

// customInt is a user-defined type based on int
type customInt int

// customPrint is a method with a value receiver
func (c customInt) customPrint() {
	fmt.Println(2 * c)
}
func main() {
	var x customInt
	x = 4 // we can also directly do like this - x := customInt(4) // direct initialization
	x.customPrint()
	(&x).customPrint()
	fmt.Println("Done")
}
