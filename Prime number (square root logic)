/* The logic for finding prime numbers in Go using the mathematical square root method involves checking whether a number is divisible by any prime numbers up to the square root of that number. Here's the logic explained step by step:

Start with an integer n greater than 1, which is the number we want to check for primality.

Iterate through all numbers i from 2 to the square root of n. This is because if n is divisible by any number greater than its square root, it would also be divisible by a smaller factor.

For each i, check if n is divisible by i without any remainder. If it is, then n is not a prime number, and we can exit the loop.

If n is not divisible by any number from 2 to the square root of n, then n is a prime number.
*/ 
package main

import (
        "fmt"
        "math"
)

func isPrime(n int)bool{
        if n<=1 {
                return false
        }
        for i:=2;i<= int(math.Sqrt(float64(n)));i++{
                if n%i == 0{
                        return false
                }
        }
        return true
}

func main(){
        for i:=100;i<=400;i++{
                if isPrime(i){
                        fmt.Printf("Number %d is a prime number\n",i)
                }else{
                        fmt.Printf("Number %d is not a prime number\n",i)

                }

        }

}
