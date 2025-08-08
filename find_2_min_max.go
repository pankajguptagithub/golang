package main

import "fmt"

const (
	maxValue = 999
	minValue = -10
)

func findMinValues(arr []int) (int, int) {
	min1, min2 := maxValue, maxValue
	for _, val := range arr {
		if val <= min1 {
			min2 = min1 // capture the value of min1
			min1 = val  // now put the new min value in min1
		} else {
			if val <= min2 {
				min2 = val
			}
		}
	}
	return min1, min2
}

func findMaxValues(arr []int) (int, int) {
	max1, max2 := minValue, minValue
	for _, val := range arr {
		if val >= max1 { // check if current value is more than max1
			max2 = max1 // first, capture the value of max1 in max2
			max1 = val  // now, replace the value of max1. Now max1 has largest value and max2 has second largest value
		} else {
			if val >= max2 {
				max2 = val // in this case when current value is less than max1 but is greater than max2, the we only replace max2
			}
		}
	}
	return max1, max2
}

func main() {
	arr := []int{1, 5, 6, 8, 10, 7, 9, 99, 0, 62, -3}
	var min1, min2, max1, max2 int
	min1, min2 = findMinValues(arr)
	fmt.Println("min1,min2 values are", min1, min2)
	max1, max2 = findMaxValues(arr)
	fmt.Println("max1,max2 values are", max1, max2)
}
