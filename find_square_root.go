package main

import "fmt"

func main() {

	var match bool
	match = false
	var result int
	m := make(map[int]int)
	n := 25
	for i := 1; i <= n; i++ {
		x := n % i
		if x == 0 {
			y := n / i
			m[i] = y
		}
	}
	fmt.Printf("factors of %d are as follows\n", n)
	for k, v := range m {
		fmt.Println(k, v)
		if k == v {
			result = k
			match = true
		}
	}
	if match == true {
		fmt.Printf("Square root of %d is %d\n", n, result)
	} else {
		fmt.Printf("%d is not a perfect square\n", n)
	}
}
