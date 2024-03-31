/*

There will be two arrays of integers. Determine all integers that satisfy the following two conditions:

The elements of the first array are all factors of the integer being considered
The integer being considered is a factor of all elements of the second array
These numbers are referred to as being between the two arrays. Determine how many such numbers exist.

Example


There are two numbers between the arrays:  and .
, ,  and  for the first value.
,  and ,  for the second value. Return .

Function Description

Complete the getTotalX function in the editor below. It should return the number of integers that are betwen the sets.

getTotalX has the following parameter(s):

int a[n]: an array of integers
int b[m]: an array of integers
Returns

int: the number of integers that are between the sets
Input Format

The first line contains two space-separated integers,  and , the number of elements in arrays  and .
The second line contains  distinct space-separated integers  where .
The third line contains  distinct space-separated integers  where .

Constraints

Sample Input

2 3
2 4
16 32 96
Sample Output

3
Explanation

2 and 4 divide evenly into 4, 8, 12 and 16.
4, 8 and 16 divide evenly into 16, 32, 96.

4, 8 and 16 are the only three numbers for which each element of a is a factor and each is a factor of all elements of b.

*/

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func getMax(a []int32)(int32){
    var max int32
    max = 0
    for i:=0;i<len(a);i++{
        if a[i] >= max{
            max = a[i]
        }
    }
    return max
}

func getTotalX(a []int32, b []int32) int32 {
    // Write your code here

    
    list1 := []int32{}

    
    //for i:=2;i<=int(max);i++{
    for i:=1;int32(i)<=getMax(b);i++{
        fmt.Println("Considering the number",i)
        fmt.Println("Considering the number against each element of array a")
        flag1 := true
        for j:=0;j<len(a);j++{
           fmt.Printf("Checking i=%d a[j]=%d \n",i,a[j]) 
           if (int32(i)%a[j]) != 0{     
               flag1 = false
               break
           }
        }   
        if flag1 == true{
            fmt.Printf("Checking %d against array b\n",i)
            flag2 := true
            for k:=0;k<len(b);k++{
                x := int32(i)                   
                fmt.Printf("Checking %d,%d\n",b[k],int32(i))
                if b[k] % x != 0{
                    fmt.Println("Inside second if")
                    flag2 = false
                    break
                }                                      
            }
            if flag2 == true{     
                    fmt.Println("Both conditions satisfied for",i)               
                    list1 = append(list1,int32(i))
            }                
        }                   
    }        
    fmt.Println(list1)
    return int32(len(list1))
}
