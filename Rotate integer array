/*Given an integer array, rotate the array to the right by k steps, where k is non negative
// ivwq 
nums=[1,2,3,4,5,6,7], k=3
output=[5,6,7,1,2,3,4]*/

// in this case length = 7 , k=3 i.e 4th index
// move elements from 0 to 3rd index to 4th index onwards
// split array in 2 parts - before k and after k
// k=1
// output=[7,1,2,3,4,5,6] current is  [2 3 4 5 6 7 1]
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	target_index := 6
	slice1 := []int{}
	slice2 := []int{}
	slice3 := []int{}
	for i := 0; i < target_index; i++ {
		slice1 = append(slice1, nums[i])
	}
	for i := target_index; i < len(nums); i++ {
		slice2 = append(slice2, nums[i])
	}
	//fmt.Println("Slice 1 is ", slice1)
	//fmt.Println("Slice 2 is ", slice2)
	for i := 0; i < len(slice2); i++ {
		slice3 = append(slice3, slice2[i])
	}
	for i := 0; i < len(slice1); i++ {
		slice3 = append(slice3, slice1[i])
	}
	fmt.Println("Final output is ", slice3)
}
