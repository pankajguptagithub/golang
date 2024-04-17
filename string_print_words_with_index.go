// Input "Hello world this is India"
// Output 1: Hello 2: world 3: this 4: is 5: India

package main

import "fmt"

func main() {
	m := make(map[int]string)
	index := 1
	s := "Hello world this is India"
	var str string
	for i := 0; i < len(s); i++ {
		if string(s[i]) != string(' ') {
			str += string(s[i])
			if i == len(s)-1 {
				m[index] = str
			}
		} else {
			m[index] = str
			index++
			str = ""
		}
	}
	fmt.Println("Output string is ", m)
}
