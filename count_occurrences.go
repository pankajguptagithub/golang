package main

import (
	"fmt"
)

func countOccurrences(str string)(map[rune]int){

	occurrences := make(map[rune]int,len(str))
	for _,v := range str{
		occurrences[v]++
	}
	return occurrences
}

func main(){
	str := "Test string"
	m := countOccurrences(str)
	fmt.Println("Count of each character")
	for k,v := range m{
		fmt.Println(string(k),v)
	}
}
