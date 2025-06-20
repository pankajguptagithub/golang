/*
We define subsequence as any subset of an array. We define a subarray as a contiguous subsequence in an array.

Given an array, find the maximum possible sum among:

all nonempty subarrays.
all nonempty subsequences.
Print the two values as space-separated integers on one line.

Note that empty subarrays/subsequences should not be considered.

Example

The maximum subarray sum is comprised of elements at inidices . Their sum is . The maximum subsequence sum is comprised of elements at indices  and their sum is .

Function Description

Complete the maxSubarray function in the editor below.

maxSubarray has the following parameter(s):

int arr[n]: an array of integers
Returns

int[2]: the maximum subarray and subsequence sums
Input Format

The first line of input contains a single integer , the number of test cases.

The first line of each test case contains a single integer .
The second line contains  space-separated integers  where .

Constraints

The subarray and subsequences you consider should have at least one element.

Sample Input 0

2
4
1 2 3 4
6
2 -1 2 3 4 -5
Sample Output 0

10 10
10 11
Explanation 0

In the first case: The maximum sum for both types of subsequences is just the sum of all the elements since they are all positive.

In the second case: The subarray  is the subarray with the maximum sum, and  is the subsequence with the maximum sum.

Sample Input 1

1
5
-2 -3 -1 -4 -6
Sample Output 1

-1 -1
Explanation 1

Since all of the numbers are negative, both the maximum subarray and maximum subsequence sums are made up of one element, .

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
 * Complete the 'maxSubarray' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
 
func max(a int32, b int32)int32{
    if a >= b{
        return a
    }else{
        return b
    }
} 

// solved with help from internet

func maxSubarray(arr []int32) []int32 {
    // Write your code here
    var maxSubArraySumHere, maxSoFar, maxSubSeqSum int32
    maxSubArraySumHere,maxSoFar,maxSubSeqSum = 0,0,0
    resultArray := []int32{}
    var allNegFlag bool 
    allNegFlag = true
    maxNegative := arr[0]
    
    for i:=0;i<len(arr);i++{
        if arr[i]>=0 {
            allNegFlag = false
            break
        }
        if arr[i] > maxNegative{
            maxNegative = arr[i]
        }
    }    
    if allNegFlag {
        maxSubSeqSum,maxSoFar = maxNegative,maxNegative        
    }else{
        for i:=0;i<len(arr);i++{      
            maxSubArraySumHere = max(arr[i],maxSubArraySumHere+arr[i])
            maxSoFar = max(maxSoFar,maxSubArraySumHere)     
            if arr[i] >= 0 {
                maxSubSeqSum += arr[i]
            }   
        }
    }
    resultArray = append(resultArray, maxSoFar)
    resultArray = append(resultArray, maxSubSeqSum)
    return resultArray
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n := int32(nTemp)

        arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        var arr []int32

        for i := 0; i < int(n); i++ {
            arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arr = append(arr, arrItem)
        }

        result := maxSubarray(arr)

        for i, resultItem := range result {
            fmt.Fprintf(writer, "%d", resultItem)

            if i != len(result) - 1 {
                fmt.Fprintf(writer, " ")
            }
        }

        fmt.Fprintf(writer, "\n")
    }

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
