/*
// Say I wanted to analyse the past week's temperatures and find the biggest rise (not just difference) in temperature.
// How would you write a program to find the biggest continuous rise in temperature without any drop in between?
//
// For example,
//
// Sun Mon    Tue  Wed   Thu   Fri   Sat
// 61   72    62   67     77    78   61   
//
// The biggest rise in temperature is 16 degrees (from Tuesday to Friday)


*/

package main

import "fmt"

func main() {
	var current_diff int
	current_diff = 0
	var max_diff int
	max_diff = 0
	//arr := []int{61, 72, 62, 67, 77, 78, 61}
	arr := []int{81, 69, 74, 79, 70, 78, 85}
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			current_diff += arr[i] - arr[i-1]
			if current_diff > max_diff {
				max_diff = current_diff
			}
		} else {
			current_diff = 0
		}
	}
	fmt.Println("max rise in tempreature is", max_diff)
}
