// An example of aliasing the fmt package as f.
package main

import (
	f "fmt"

	g "github.com/yuanyewa/gofullstack/volume1/section2/greetingspackage"
)

func main() {

	f.Println("Hello Gopher!")
	g.PrintGreetings()

}
