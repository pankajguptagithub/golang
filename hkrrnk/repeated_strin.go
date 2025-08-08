/*
There is a string, , of lowercase English letters that is repeated infinitely many times. Given an integer, , find and print the number of letter a's in the first  letters of the infinite string.

Example


The substring we consider is , the first  characters of the infinite string. There are  occurrences of a in the substring.

Function Description

Complete the repeatedString function in the editor below.

repeatedString has the following parameter(s):

s: a string to repeat
n: the number of characters to consider
Returns

int: the frequency of a in the substring
Input Format

The first line contains a single string, .
The second line contains an integer, .

Constraints

For  of the test cases, .
Sample Input

Sample Input 0

aba
10
Sample Output 0

7
Explanation 0
The first  letters of the infinite string are abaabaabaa. Because there are  a's, we return .

Sample Input 1

a
1000000000000
Sample Output 1

1000000000000
Explanation 1
Because all of the first  letters of the infinite string are a, we return .
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
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
    // Write your code here
    //var count int64
    //count = 0
    //var actualString []string 
    var initialCount int64 = 0
    var actualCount int64 = 0
    
    if n>=int64(len(s)){
        for i:=0;i<len(s);i++{
        if string(s[i]) == "a"{
            initialCount ++
            }
         }
         rep := n/int64(len(s)) // how many times n is larger than given string 
         fmt.Printf("For %d and %s, value of rep is %d",n,s,rep)
         actualCount = actualCount + rep*int64(initialCount)
         charRemaining := n%int64(len(s))
         for i:=0;i<int(charRemaining);i++{
            if string(s[i]) == "a"{
            actualCount++
            }      
         }
    }else{
        for i:=0;int64(i)<n;i++{
        if string(s[i]) == "a"{
            initialCount ++
            }
         }
         actualCount = initialCount
    } 
    fmt.Printf("current count in %s is %d\n", s,initialCount )
    return int64(actualCount)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    result := repeatedString(s, n)

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
