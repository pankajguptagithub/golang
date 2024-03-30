//find in an array, pair of numbers whose sum is equal to a given number
package main

import (
	"fmt"
)
func find_pairs(arr []int, target int)(map[int]int){
	m := make(map[int]bool,len(arr))
	pair := make(map[int]int)

	for i:=0; i<len(arr);i++{
		m[arr[i]] = false 
	}
	fmt.Println("Initial map is ",m)
	for i:=0; i<len(arr);i++{
		diff := target - arr[i]
		m[arr[i]] = true
		if m[diff]{
			pair[arr[i]] = diff
		}
	}
	return pair

}
func main(){
	arr := []int{1,3,2,4,8,9,7,5,6}
	target := 11
	m := make(map[int]int)
	m = find_pairs(arr,target)
	fmt.Println("Displaying the pairs")
	for k,v := range m{
		fmt.Println(k,v)
	}
}
