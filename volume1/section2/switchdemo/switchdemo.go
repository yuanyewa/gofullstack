// An example of declaring a switch statement.
package main

import "fmt"

func main() {

	z := 36

	// A switch statement example
	switch z {
	case 1:
		fmt.Println("z is equal to 1")

	case 2:
		fmt.Println("z is equal to 2")

	case 3:
		fmt.Println("z is equal to 3")

	default:
		fmt.Println("z does not equal 1, 2, or 3")

	}
	var (
		x int
		y string
	)
	switch interface{} (x).(type) {
	case int:
		fmt.Println("I'm int")
	default:
		fmt.Println("I'm not int")
	}
	switch interface{} (y).(type) {
	case int:
		fmt.Println("I'm int")
	case string:
		fmt.Println("I'm string")
	default:
		fmt.Println("I'm not int")
	}
}
