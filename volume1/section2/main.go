package main

import (
	"fmt"

	"github.com/yuanyewa/gofullstack/volume1/section2/greetingspackage"
)

func main() {
	fmt.Println(greetingspackage.MagicNumber)
	fmt.Println(greetingspackage.MyVar)
	greetingspackage.PrintGreetings()
}
