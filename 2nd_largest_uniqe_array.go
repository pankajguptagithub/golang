package main

import "fmt"

const (
	min = -1
)

func findUniq(arr []int) map[int]bool {
	uniqElementsMap := make(map[int]bool)
	for _, val := range arr {
		if uniqElementsMap[val] {
			uniqElementsMap[val] = false
		} else {
			uniqElementsMap[val] = true
			
		}
	}
	return uniqElementsMap
}
func findMax(m map[int]bool) int {
	max1, max2 := min, min
	for val, flag := range m {
		if flag == true {
			if val >= max1 {
				max2 = max1
				max1 = val
			} else {
				if val >= max2 {
					max2 = val
				}
			}
		}
	}
	return max2
}

func main() {
	//arr := []int{10, 4, 3, 50, 23, 50, 3, 90}
	arr := []int{5, 5, 5, 5, 5}
	uniqElements := findUniq(arr)
	fmt.Println("unique slice is ", uniqElements)
	max2 := findMax(uniqElements)
	fmt.Println("Second largest unique element in the array is ", max2)

}
