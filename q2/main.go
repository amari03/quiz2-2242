// channels are a way that go routines can communicate with each other
//can send data from one go routine to another

//multiplication results example
//In this example we have dirctional channels where data is sent and recieved.

package main

import (
	"fmt"
)

// This function now accepts a channel as its second argument...
func multiplyByTwo(num int, result chan<- int) {
	//result chan<- int ... this tells you that data can only be sent into the channel
	received := num * 3

	//... and pipes the result into it
	result <- received
}

func main() {
	n := 4

	// This is where we "make" the channel, which can be used
	// to move the `int` datatype
	result := make(chan int)

	// We still run this function as a goroutine, but this time,
	// the channel that we made is also provided
	go multiplyByTwo(n, result)

	// Once any output is received on this channel, print it to the console and proceed
	fmt.Println(<-result)
}
