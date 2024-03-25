/*

Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.

Example

The minimum sum is  and the maximum sum is . The function prints

16 24
Function Description

Complete the miniMaxSum function in the editor below.

miniMaxSum has the following parameter(s):

arr: an array of  integers
Print

Print two space-separated integers on one line: the minimum sum and the maximum sum of  of  elements.

Input Format

A single line of five space-separated integers.

Constraints


Output Format

Print two space-separated long integers denoting the respective minimum and maximum values that can be calculated by summing exactly four of the five integers. (The output can be greater than a 32 bit integer.)

Sample Input

1 2 3 4 5
Sample Output

10 14
Explanation

The numbers are , , , , and . Calculate the following sums using four of the five integers:

Sum everything except , the sum is .
Sum everything except , the sum is .
Sum everything except , the sum is .
Sum everything except , the sum is .
Sum everything except , the sum is .
Hints: Beware of integer overflow! Use 64-bit Integer.

Need help to get started? Try the Solve Me First problem

*/


package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

/*func miniMaxSum(arr []int64) {
    var sum,min,max int64 
    sum = 0
    min = math.Maxint64    
    max = math.Minint64  
    for i:=0;i<len(arr);i++{
        sum += arr[i]
        if arr[i] < min{
            min = arr[i]
        }
        if arr[i] > max{
            max = arr[i]
        }    
    }
    fmt.Println("Total sum is ",sum)
    fmt.Println("Max is ",max)
    fmt.Println("Min is ",min)
    min_sum := sum - max
    max_sum := sum - min    
    fmt.Println("Minimum sum is",min_sum)
    fmt.Println("Maximum sum is ",max_sum)
   
}*/
func miniMaxSum(arr []int64) {
    var max_sum int64
    var min_sum int64
    var total_sum int64
    
    for i:=0;i<len(arr);i++{
       total_sum += arr[i]
    }    
    min_sum = math.MaxInt64
    max_sum = math.MinInt64
    for i:=0;i<len(arr);i++{
        current_sum := total_sum - arr[i]
        if current_sum < min_sum{
            min_sum = current_sum
        }
        if current_sum > max_sum{
            max_sum = current_sum
        }        
    }        
    fmt.Println("Min sum is",min_sum)
    fmt.Println("Max sum is",max_sum)
}
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int64

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int64(arrItemTemp)
        arr = append(arr, arrItem)
    }
    miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

/*
Adding another solution here 
*/

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func bubble_sort(arr []int64){
    for i:=0;i<len(arr);i++{
        for j:=0;j+i<len(arr)-1;j++{
            if arr[j] > arr[j+1]{
                temp := arr[j+1]
                arr[j+1] = arr[j]
                arr[j] = temp
            }
        }
    }      
}
func miniMaxSum(arr []int64) {
    var min_sum, max_sum int64
    min_sum = 0
    max_sum = 0
    bubble_sort(arr)
    for i:=0;i<len(arr)-1;i++{
        min_sum += arr[i]
    }
    for i:=1;i<len(arr);i++{
        max_sum += arr[i]
    }
    //min_sum = arr[0]+arr[1]+arr[2]+arr[3]
    //max_sum = arr[1]+arr[2]+arr[3]+arr[4]
    fmt.Println(min_sum,max_sum)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int64

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int64(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
