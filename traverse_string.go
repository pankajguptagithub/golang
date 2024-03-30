package main 

import (
	"fmt"
)

func main(){
	s := "Hello World"
	for k,v := range s{
		fmt.Println(k,string(v))
	}
	for _,v := range s{
		fmt.Println(string(v))
	}
	for k,_ := range s{
		fmt.Println(k)
	}
	s2 := []string{"Hello World2"}
	fmt.Println(s2)
	for k,v := range s2{
		fmt.Println(k,string(v))
	}
}
