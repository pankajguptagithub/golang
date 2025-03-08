package main 

import (
	"fmt"
)

// Integer type implementing the Stringer interface


type (
	Custom_integer int
	Stringer interface {
		Custom_string() string
	}
)	
// Custom_string returns the string representation of the Integer type


func (ci Custom_integer) Custom_string()(string){
		return fmt.Sprintf("Integer value: %d",ci)
}

func main(){
	var x Stringer
	x = Custom_integer(42)
	fmt.Println(x.Custom_string())
}


