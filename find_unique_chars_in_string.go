package main

import (
	"fmt"
)

func find_unique(str string)[]string{
	m := make(map[rune]bool)
	out := []string{}

	for _,v := range str{
		if !m[v]{
			m[v] = true
			out = append(out,string(v))
		}
	}
	return out
}
func main(){
	str := "Hello world234"
	uniq := find_unique(str)
	fmt.Println("List of unique chars in string",uniq)
}

