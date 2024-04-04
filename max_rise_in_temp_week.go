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


/*
Other attempts of the code which did not work 

package main
import "fmt"



func main(){
	var max_diff int
  max_diff = 0
  var current_diff int
  current_diff = 0
  var current_index int
  current_index = 0
  arr := []int{61,72,62,67,77,78,61}
  i:=0
  j:=0  

  for i:=0;i<len(arr);{
    for j:=i;j<len(arr);{ 
    if arr[j] > arr[i]{
      j++
    }else{
      current_diff = arr[j] - arr[i]
      if current_diff > max_diff{
        max_diff = current_diff
      }
      i++
    }
    }
  }
  fmt.Println(max_diff)   
      
 /* }
  
  //for ;i<len(arr);{
  for ;i<len(arr),j<len(arr); {
    if arr[j] >= arr[i]{
      current_diff = arr[j] - arr[i]
      if current_diff > max_diff{
        max_diff = current_diff
      }
      j++
    }else{
      i++
    }    
  }
  */
  //fmt.Println("Max diff is ",max_diff)
  /*for ;i<len(arr);{
    if arr[i+1] >= arr[i]{
      current_diff = arr[i+1] - arr[i]
      i++
    }else{
      
    }
  
  
  
  for ;i<len(arr);{
    start_index := i 
    if arr[i+1] >= arr[i]{
      fmt.Println("Inside if ")
      current_index = i+1
      fmt.Println("current index is ",current_index)
      i++
    }else{    
      fmt.Println("Inside else")
      current_diff = arr[current_index] - arr[start_index]
      fmt.Println("Current diff is ",current_diff)
      if current_diff >= max_diff{
        fmt.Println("Inside nested if")
        max_diff = current_diff
        fmt.Println("Max diff now is ",max_diff)
      }  
      i++
    }           
  }
fmt.Println(max_diff)  
}
/*

*/
