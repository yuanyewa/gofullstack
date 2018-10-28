// An example of a buffered channel (asynchronous).
package main

import "fmt"

func main() {

	// We create a buffered channel of strings with a capacity of 3
	// This means the channel buffer can hold up to 3 values
	messageQueue := make(chan string, 3)
	messageQueue <- "one"
	messageQueue <- "two"
	messageQueue <- "three"

	// We drain the messageQueue by receiving all the values from the
	// buffered channel.
	fmt.Println(<-messageQueue)
	fmt.Println(<-messageQueue)
	fmt.Println(<-messageQueue)
	val := make(chan int) //yuanye: why buffered channel can be written directly; unbuffered must be "go"-ed
	//https://groups.google.com/forum/#!topic/golang-nuts/n1jKyGu6krk
	// for un-buffered channel, when you are writing to it, someone must be listening to it
	go func() { val <- 10 }()
	fmt.Println(<-val)
}
