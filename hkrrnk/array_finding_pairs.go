Given an array of integers and a target value, determine the number of pairs of array elements that have a difference equal to the target value.

Example


There are three values that differ by : , , and . Return .

Function Description

Complete the pairs function below.

pairs has the following parameter(s):

int k: an integer, the target difference
int arr[n]: an array of integers
Returns

int: the number of pairs that satisfy the criterion
Input Format

The first line contains two space-separated integers  and , the size of  and the target value.
The second line contains  space-separated integers of the array .

Constraints

each integer  will be unique
Sample Input

STDIN       Function
-----       --------
5 2         arr[] size n = 5, k =2
1 5 3 4 2   arr = [1, 5, 3, 4, 2]
Sample Output

3
Explanation

There are 3 pairs of integers in the set with a difference of 2: [5,3], [4,2] and [3,1]. .


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
 * Complete the 'pairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY arr
 */


//arr = [1, 5, 3, 4, 2]
//k = 1 
// learnt the technique over internet 
func pairs(k int32, arr []int32) int32 {
    // Write your code here
    var count int32 
   count = 0
   m := make(map[int32]bool)
   
   for i:=0;i<len(arr);i++{
        m[arr[i]] = true
        
        if m[arr[i]+k]{
            count++
        }
        if m[arr[i]-k]{
            count++
        }
   }
   return count 
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := pairs(k, arr)

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
