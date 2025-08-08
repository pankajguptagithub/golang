/*
Your local library needs your help! Given the expected and actual return dates for a library book, create a program that calculates the fine (if any). The fee structure is as follows:

If the book is returned on or before the expected return date, no fine will be charged (i.e.: .
If the book is returned after the expected return day but still within the same calendar month and year as the expected return date, .
If the book is returned after the expected return month but still within the same calendar year as the expected return date, the .
If the book is returned after the calendar year in which it was expected, there is a fixed fine of .
Charges are based only on the least precise measure of lateness. For example, whether a book is due January 1, 2017 or December 31, 2017, if it is returned January 1, 2018, that is a year late and the fine would be .

Example


The first values are the return date and the second are the due date. The years are the same and the months are the same. The book is  days late. Return .

Function Description

Complete the libraryFine function in the editor below.

libraryFine has the following parameter(s):

d1, m1, y1: returned date day, month and year, each an integer
d2, m2, y2: due date day, month and year, each an integer
Returns

int: the amount of the fine or  if there is none
Input Format

The first line contains  space-separated integers, , denoting the respective , , and  on which the book was returned.
The second line contains  space-separated integers, , denoting the respective , , and  on which the book was due to be returned.

Constraints

Sample Input

9 6 2015
6 6 2015
Sample Output

45
Explanation

Given the following dates:
Returned: 
Due: 

Because , we know it is less than a year late.
Because , we know it's less than a month late.
Because , we know that it was returned late (but still within the same month and year).

Per the library's fee structure, we know that our fine will be . We then print the result of  as our output.
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
 * Complete the 'libraryFine' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER d1
 *  2. INTEGER m1
 *  3. INTEGER y1
 *  4. INTEGER d2
 *  5. INTEGER m2
 *  6. INTEGER y2
 */

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
    var ret int32 = 0
    if d2 == d1 && m2 == m1 && y2 == y1{
        ret = 0
    }else if y1 > y2{
        ret=10000
    }else if y2 == y1 && m1 > m2{
        ret=500*(m1-m2)
    }else if y2 == y1 && m2 == m1 && d1 > d2{
        ret=15*(d1-d2)
    }
    return ret
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    d1Temp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    d1 := int32(d1Temp)

    m1Temp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    m1 := int32(m1Temp)

    y1Temp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
    checkError(err)
    y1 := int32(y1Temp)

    secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    d2Temp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
    checkError(err)
    d2 := int32(d2Temp)

    m2Temp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
    checkError(err)
    m2 := int32(m2Temp)

    y2Temp, err := strconv.ParseInt(secondMultipleInput[2], 10, 64)
    checkError(err)
    y2 := int32(y2Temp)

    result := libraryFine(d1, m1, y1, d2, m2, y2)

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

