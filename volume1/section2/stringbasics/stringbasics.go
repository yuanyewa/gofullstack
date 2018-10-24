// String examples covering declaration, literal definition, accessing character at value
// and concatenating strings using the + operator.
package main

import "fmt"

type myT int

func (m myT) String() string {
	return "this is meeee!"
}
func main() {

	var subject string = "Gopher"

	fmt.Println("First element of Gopher string: ", string("Gopher"[0]))
	fmt.Printf("The first value of the subject string: %v\n", string(subject[0]))
	fmt.Printf("The last value of the subject string: %v\n", string(subject[len(subject)-1]))
	fmt.Println("Hello " + subject + "!")
	fmt.Println(len("this is me"))
	fmt.Printf("sdfsf%v%v\n",3.13, myT(3))
	x := fmt.Sprintf("sdfsfsfsf%v\n", "!!!!!!!!!!!!!!!")
	fmt.Printf(x)
}
