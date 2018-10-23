// An example of using a command line flag.
package main

import (
	"flag"
	"fmt"
)

func main() {

	var gopherName, myFlag string
	//flag.StringVar syntax: address, flag, default, help msg
	flag.StringVar(&gopherName, "gophername", "Gopher", "The name of the Gopher")
	flag.StringVar(&myFlag, "myflag", "a flag", "enter your flag name")
	flag.Parse()
	fmt.Println("Hello " + gopherName + myFlag + "!")
}
