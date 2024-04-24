/*
Find the longest palindrome substring 
/ Input: str = “forteeksskeetfor” 
// Output: "teeksskeet"
// Explanation: There are several possible palindromic substrings like “kssk”, “ss”, “eeksskee” etc. But the substring “teeksskeet” is the longest among all.
// """

*/

package main

import (
	"fmt"
)

func generateSubstrings(s string) []string {
	var subStringList []string
	for i := 0; i < len(s); i++ {
		//str := string(s[i])
		for j := i; j < len(s); j++ {
			//str += string(s[j])
			//fmt.Println("Sub string generated is ", str)
			str := s[i:j]
			if str != "" {
				subStringList = append(subStringList, s[i:j])
			}
		}
	}
	return subStringList
}
func checkPalindrome(s string) bool {
	fmt.Println("Checking string for palindrome", s)
	var flag bool
	flag = true
	for i := 0; i < len(s)/2; i++ {
		if string(s[i]) == string(s[len(s)-i-1]) {
			continue
		} else {
			flag = false
			break
		}
	}
	return flag
}
func main() {
	fmt.Println("Hello")
	var listSubstrings []string
	var result string
	max_length := 0
	listSubstrings = generateSubstrings("forteeksskeetfor")
	for i := 0; i < len(listSubstrings); i++ {
		if checkPalindrome(listSubstrings[i]) {
			fmt.Println("String is palindrome and the length is - ", listSubstrings[i], len(listSubstrings[i]))
			if len(listSubstrings[i]) > max_length {
				max_length = len(listSubstrings[i])
				fmt.Println("Max length is ", len(listSubstrings[i]))
				result = listSubstrings[i]
			}
		}
	}
	fmt.Println("Max substring is ", result)
}
