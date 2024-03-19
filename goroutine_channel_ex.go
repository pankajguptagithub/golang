package main

import (
	"fmt"
)

func main() {
	c := make(chan int) // Allocate a channel.

	// Start the sort in a goroutine
	go func() {
		fmt.Println("Hello, world")
		c <- 1 // Send a signal
		fmt.Println("Goroutine Finished.")
	}()
	//<-c // Receive a value from the channel
	fmt.Println("Program Completed!")
}

/*
Output 
Program Completed!

Here this program doesn't print what's inside go routine because main go routine would not wait for the goroutine to complete. 
Main function will print line 17 and exit. 
However, in the given code if we just uncomment line 16, the problem get's solved because we are writing something into channel and 
receiving it in main so overall it works 
*/
// Another variant of the code 

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // Allocate a channel.
	// Start the sort in a goroutine
	go func() {
		fmt.Println("Hello, world")
		c <- 1 // Send a signal
		fmt.Println("Goroutine Finished.")

	}()
	time.Sleep(1 * time.Second)
	fmt.Println("Program Completed!")
}

/*
Hello, world
Program Completed!

Here the line number 44 doesn't get printed because of line no 43. 

1. we are sending data into channel but not receiving it anywhere. So if we interchance line 43 and 44, everything will get printed. 
2. Or if we add the line "<- c" in the main code then also it works. channel operation (c <- 1) is blocking and the main goroutine isn't reading from the channel
   This is a blocking operation on an unbuffered channel because there's no buffer space to store the value. As a result, this operation blocks until another goroutine is ready to receive from the channel.

To make the channel operation non-blocking, you can use a buffered channel or utilize Go's select statement. Here's how you can modify the code to make the channel operation non-blocking using a buffered channel:

1  func main() {
2      c := make(chan int, 1) // Allocate a buffered channel with capacity 1.
3      // Start the sort in a goroutine
4      go func() {
5          fmt.Println("Hello, world")
6          select {
7          case c <- 1: // Try to send a signal into the channel, but don't block if the channel is full.
8          default:
9              fmt.Println("Channel is full. Cannot send.")
10         }
11         fmt.Println("Goroutine Finished.")
12     }()
13     time.Sleep(1 * time.Second)
14     fmt.Println("Program Completed!")
15 }
In this modified version:

Line 2: The channel c is created as a buffered channel with capacity 1. 
This means it can hold one value without blocking a sender.
Inside the goroutine, a select statement is used to attempt sending 
the value 1 into the channel c. However, if the channel is full and 
cannot receive the value immediately, it will execute the default 
case and print "Channel is full. Cannot send." without blocking.
This modification ensures that the send operation does not block 
the goroutine, even if the channel is not ready to receive the 
value immediately. It provides non-blocking behavior by handling 
the case where the channel is full.

*/
