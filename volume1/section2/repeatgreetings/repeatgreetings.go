// An example of using a command line argument.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var repeatCount int
	var err error

	var f string

	flag.StringVar(&f, "flag", "aFlag", "enterAFlag")
	flag.Parse()
	if len(os.Args) >= 2 {
		repeatCount, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < repeatCount; i++ {
			fmt.Println("Hello Gopher!", f)
		}

	}

	fmt.Println("FFS: how to combine flag and os.Args")

}
