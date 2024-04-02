package main

import "fmt"

func main() {
	arr := []int{100,5, 10, 101, 85, 18, 75, 36, 9, 93, 82, 84, 74, 90, 76, 79, 80, 89}
	max := 0
	prev_max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			prev_max = max
			max = arr[i]
		} else {
			if arr[i] > prev_max {
				prev_max = arr[i]
			}
		}
	}
	fmt.Println("Max element is ", max)	
	fmt.Println("2nd largest element is ", prev_max)
}
