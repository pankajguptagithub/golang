package main

import (
	"fmt"
)

func read(ch chan int) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan int)
	go read(ch) // interchanging this line with below one would cause deadlock
	ch <- 3
	fmt.Println("Done")
}
