/* 
Given an array of integers, find the longest subarray where the absolute difference between any two elements is less than or equal to .

Example


There are two subarrays meeting the criterion:  and . The maximum length subarray has  elements.

Function Description

Complete the pickingNumbers function in the editor below.

pickingNumbers has the following parameter(s):

int a[n]: an array of integers
Returns

int: the length of the longest subarray that meets the criterion
Input Format

The first line contains a single integer , the size of the array .
The second line contains  space-separated integers, each an .

Constraints

The answer will be .
Sample Input 0

6
4 6 5 3 3 1
Sample Output 0

3
Explanation 0

We choose the following multiset of integers from the array: . Each pair in the multiset has an absolute difference  (i.e.,  and ), so we print the number of chosen integers, , as our answer.

Sample Input 1

6
1 2 2 3 1 2
Sample Output 1

5
Explanation 1

We choose the following multiset of integers from the array: . Each pair in the multiset has an absolute difference  (i.e., , , and ), so we print the number of chosen integers, , as our answer.

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

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

// previous unsuccessful attempt
/*
//var i,j float32
   
    max_length := 0
 for i:=0;i<len(a);i++{
        fmt.Println("============")
        fmt.Println("outer array element is ",a[i])
        temp_arr := []int32{}
        temp_arr = append(temp_arr,a[i])
        for j:=i+1;j<len(a);j++{
            fmt.Println("inner array element is ",a[j])
            if int32(math.Abs(float64(a[i]) - float64(a[j]))) <=1 {
                fmt.Printf("Inside if, appending %d\n",a[j])
                temp_arr = append(temp_arr,a[j])   
            }
         }
         fmt.Printf("for array element %d, temp_arr is %d\n",a[i],temp_arr)
         if len(temp_arr) >= max_length{
            max_length = len(temp_arr)
         }
    }
    return int32(max_length)
*/

func pickingNumbers(a []int32) int32 {
    // Write your code here
    count := make(map[int32]int32)
    var longestSubArray int32
    longestSubArray = 0 
    for _,num := range a{
        count[num]++
                
        if count[num] >= longestSubArray{
            longestSubArray = count[num]
        }        
        if count[num]+count[num-1] >= longestSubArray{
            longestSubArray = count[num]+count[num-1]
        }
        if count[num]+count[num+1] >= longestSubArray{
            longestSubArray = count[num]+count[num+1]
        }        
    }
    return longestSubArray
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var a []int32

    for i := 0; i < int(n); i++ {
        aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
        checkError(err)
        aItem := int32(aItemTemp)
        a = append(a, aItem)
    }

    result := pickingNumbers(a)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
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


