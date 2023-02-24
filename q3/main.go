// create and use Goroutines
// Goroutines are functions or methods that run concurrently with other functions or methods.
// In Go, we use goroutines to create concurrent programs.
// Concurrent programs are able to run multiple processes at the same time.
// we use goroutines that help us to achieve concurrency in programming.
package main

import (
	"fmt"
	"time"
)

func message() {
	fmt.Println("Howdy Partner! This is the goroutine.")
}

// main function runs it own go routine
func main() {

	// Calling Goroutine
	go message()

	time.Sleep(1 * time.Second)
	// Calling normal function
	fmt.Println("Welcome from goroutine main")

}

/*
The sleep method is just a temporary fix.
*/
