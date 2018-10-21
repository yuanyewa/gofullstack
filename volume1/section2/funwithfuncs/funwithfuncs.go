// Various examples of functions in Go.
package main

import (
	"fmt"
)

// An example of a function that takes one input parameter, but does not return
// anything back. Thus, it is considered a void function.

func oddOrEven(x int) {

	// Functions have a local scope. The variable x is only accessible within this function.
	// For that reason we can consider x to be a local variable of the function.
	if x%2 == 0 {
		fmt.Println("The number,", x, ",is even.")
	} else {
		fmt.Println("The number,", x, ",is odd.")
	}

	// The variable y is visible inside this function, because it exists in the global
	// scope. Therefore, y is a global variable.
	if y%2 == 0 {
		fmt.Println("The number,", y, ",is even.")
	} else {
		fmt.Println("The number,", y, ",is odd.")
	}
}

var y = 4

// named return parameters are already defined at top level. Can use directly
// function with variable number of input
func myfunc(xx ...interface{}) (x, y int) {
	oddOrEven(5)
	x = 4
	y = 5
	for _, v := range xx { // range function to loop over
		fmt.Println(v)
	}
	return //naked return; only use for short functions
}

// An example of a function that takes 2 input parameters and returns an output
// of type int
func sumTwoNumbers(x int, y int) int {

	return x + y
}

// An example of a function returning multiple, named output parameters
func sumAndDiffOperationsOnTwoNumbers(x, y int) (sum int, difference int) {
	return x + y, x - y
}

// An example of a variadic function, where we can supply a varying number of
// inputs of the same type
func multiSum(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}

	return sum
}

// main() is a niladic function because it doesn't accept any arguments
func main() {
	xx, _ := myfunc([]interface{}{10, "a"}...)
	fmt.Println(xx)
	var numberToCheck = 7
	oddOrEven(numberToCheck)

	fmt.Println("The sum of 100 and 8: ", sumTwoNumbers(100, 8))

	sum, difference := sumAndDiffOperationsOnTwoNumbers(100, 8)
	fmt.Printf("Input: 100 and 8, The Sum: %d Difference: %+v\n", sum, difference)

	fmt.Println("Multi Sum result: ", multiSum(0, 5, 76, 2, 3, 4, 8, 1, 9))

	// Example of an anonymous function
	func() {
		fmt.Println("Hi I'm an anonymous function. They call me that, because I don't have a name!")
	}()

}
