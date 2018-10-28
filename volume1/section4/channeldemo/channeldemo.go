// An example of sending and receiving a value through a channel.
package main

import (
	"fmt"
)

func writeMessageToChannel(message chan string) {
	message <- "Hello Gopher!"

}

func main() {

	fmt.Println("Channel Demo")

	message := make(chan string)
	defer close(message)
	go writeMessageToChannel(message)

	val := make(chan int, 2)
	go func() { val <- 10; val <- 20 }()
	fmt.Println(<-val, <-val)
	fmt.Println("Greeting from the message channel: ", <-message)
}
