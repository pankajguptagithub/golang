/*
Given an array of integers, where all elements but one occur twice, find the unique element.
Example
The unique element is .
Function Description
Complete the lonelyinteger function in the editor below.
lonelyinteger has the following parameter(s):
int a[n]: an array of integers
Returns
int: the element that occurs only once
Input Format
The first line contains a single integer, , the number of integers in the array.
The second line contains  space-separated integers that describe the values in .
Constraints


It is guaranteed that  is an odd number and that there is one unique element.
, where .
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
 * Complete the 'lonelyinteger' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */


/*func lonelyinteger(a []int32) int32 {
    var ans int32 
    for i:=0;i<len(a);i++{
            count := 0
            for j:=0;j<len(a);j++{
               if a[i] == a[j]{
                    count++                    
               }
            }
        fmt.Println("Element is ",a[i])
        fmt.Println("Element count is ",count)    
        if count == 1{
            ans = a[i]        
        }
    }
    return ans
}
*/
/* func lonelyinteger(a []int32) int32 {
    var return_var int32
    counts := make(map[int32]int)
    
    for _, num := range a{
        counts[num]++        
    }
    
    for num, count := range counts{
        
        if count == 1{
            return_var = num
        }
        
    }
    return return_var
}*/


func lonelyinteger(a []int32) int32 {
    var result int32
    for _, num := range a{
        result = result ^ num        
    }
    return result
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


    result := lonelyinteger(a)


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
