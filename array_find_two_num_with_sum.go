package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	m := make(map[int]bool)
	for _, val := range arr {
		m[val] = true
		compliment := target - val
		if m[compliment] {
			fmt.Println("Found ", val, compliment)
			break
		}
	}
}
