/*
Sam's house has an apple tree and an orange tree that yield an abundance of fruit. Using the information given below, determine the number of apples and oranges that land on Sam's house.

In the diagram below:

The red region denotes the house, where  is the start point, and  is the endpoint. The apple tree is to the left of the house, and the orange tree is to its right.
Assume the trees are located on a single point, where the apple tree is at point , and the orange tree is at point .
When a fruit falls from its tree, it lands  units of distance from its tree of origin along the -axis. *A negative value of  means the fruit fell  units to the tree's left, and a positive value of  means it falls  units to the tree's right. *
Apple and orange(2).png

Given the value of  for  apples and  oranges, determine how many apples and oranges will fall on Sam's house (i.e., in the inclusive range )?

For example, Sam's house is between  and . The apple tree is located at  and the orange at . There are  apples and  oranges. Apples are thrown  units distance from , and  units distance. Adding each apple distance to the position of the tree, they land at . Oranges land at . One apple and two oranges land in the inclusive range  so we print

1
2
Function Description

Complete the countApplesAndOranges function in the editor below. It should print the number of apples and oranges that land on Sam's house, each on a separate line.

countApplesAndOranges has the following parameter(s):

s: integer, starting point of Sam's house location.
t: integer, ending location of Sam's house location.
a: integer, location of the Apple tree.
b: integer, location of the Orange tree.
apples: integer array, distances at which each apple falls from the tree.
oranges: integer array, distances at which each orange falls from the tree.
Input Format

The first line contains two space-separated integers denoting the respective values of  and .
The second line contains two space-separated integers denoting the respective values of  and .
The third line contains two space-separated integers denoting the respective values of  and .
The fourth line contains  space-separated integers denoting the respective distances that each apple falls from point .
The fifth line contains  space-separated integers denoting the respective distances that each orange falls from point .

Constraints

Output Format

Print two integers on two different lines:

The first integer: the number of apples that fall on Sam's house.
The second integer: the number of oranges that fall on Sam's house.
Sample Input 0

7 11
5 15
3 2
-2 2 1
5 -6
Sample Output 0

1
1
Explanation 0

The first apple falls at position .
The second apple falls at position .
The third apple falls at position .
The first orange falls at position .
The second orange falls at position .
Only one fruit (the second apple) falls within the region between  and , so we print  as our first line of output.
Only the second orange falls within the region between  and , so we print  as our second line of output.

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
 * Complete the 'countApplesAndOranges' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER s
 *  2. INTEGER t
 *  3. INTEGER a
 *  4. INTEGER b
 *  5. INTEGER_ARRAY apples
 *  6. INTEGER_ARRAY oranges
 */

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
    // Write your code here
    countApples := 0
    countOranges := 0 
    for _, apple := range apples {
        applePosition := a + apple
        if applePosition >= s && applePosition <= t {
            countApples++
        }        
    }
    for _, orange := range oranges {
        orangePosition := b + orange
        if orangePosition >= s && orangePosition <= t {
            countOranges++
        }        
    }
    fmt.Printf("%d\n%d",countApples, countOranges)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    sTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    s := int32(sTemp)

    tTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    t := int32(tTemp)

    secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    aTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
    checkError(err)
    a := int32(aTemp)

    bTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
    checkError(err)
    b := int32(bTemp)

    thirdMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    mTemp, err := strconv.ParseInt(thirdMultipleInput[0], 10, 64)
    checkError(err)
    m := int32(mTemp)

    nTemp, err := strconv.ParseInt(thirdMultipleInput[1], 10, 64)
    checkError(err)
    n := int32(nTemp)

    applesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var apples []int32

    for i := 0; i < int(m); i++ {
        applesItemTemp, err := strconv.ParseInt(applesTemp[i], 10, 64)
        checkError(err)
        applesItem := int32(applesItemTemp)
        apples = append(apples, applesItem)
    }

    orangesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var oranges []int32

    for i := 0; i < int(n); i++ {
        orangesItemTemp, err := strconv.ParseInt(orangesTemp[i], 10, 64)
        checkError(err)
        orangesItem := int32(orangesItemTemp)
        oranges = append(oranges, orangesItem)
    }

    countApplesAndOranges(s, t, a, b, apples, oranges)
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
Above logic which was accepted by Hackerrank was taken from chat gpt after my logic wasn't accepted
Here's the logic I wrote 

count_apples := 0
    count_oranges := 0 
    for i:=0;i<len(apples);i++{
        fmt.Println("Value of apples is ",apples[i])
        fmt.Println("Value of a is ",a)
        fmt.Println("Its position is ",apples[i]+a)
        if (apples[i] + a) >= s && (apples[i] + a) <= t{
            count_apples++
        }        
    }
     for i:=0;i<len(oranges);i++{
        fmt.Println("Value of oranges is ",oranges[i])
        fmt.Println("Value of b is ",b)
        fmt.Println("Its position is",oranges[i]+b)
        if (oranges[i] + b) >= s && (oranges[i] + b) <= t{
            count_oranges++
        }        
    }
    fmt.Printf("%d\n%d",count_apples,count_oranges)

*/
