// An example of an empty interface, and some useful things we can do with it.
package main

import (
	"fmt"
)

// step 1: define an interface type, that specifies a function
type myShape interface {
	Area() float32
}

// step 2: define a function with input of interface, and uses the interface function
func getArea(m myShape) float32 {
	return m.Area()
}

// step 3: define another type
type triangle struct {
	base   float32
	height float32
}

// step 4: implement the interface function for the type
func (t triangle) Area() float32 {
	return t.base * t.height / 2
}

// now, repeat step 3 for another type
type circle struct {
	radius float32
}

func (c circle) Area() float32 {
	return 2 * 3.14 * c.radius * c.radius
}
func giveMeARandomShape() interface{} {

	var shape interface{}
	randomShapesSlice := make([]interface{}, 2)

	tri := triangle{base: 3, height: 4}
	cir := circle{1}

	// Let's store the shapes into a slice data structure
	randomShapesSlice[0] = tri
	randomShapesSlice[1] = cir
	shape = randomShapesSlice[1]

	return shape
}

func main() {

	// empty interface allows combine different types together
	shapesSlice := []interface{}{triangle{3, 4}, circle{2}}
	for index, shape := range shapesSlice {
		fmt.Printf("Shape in index, %d, of the shapesSlice is a  %T\n", index, shape)
	}
	fmt.Println("\n")

	aRandomShape := giveMeARandomShape()
	fmt.Printf("The type of aRandomShape is %T\n", aRandomShape)
	switch t := aRandomShape.(type) { // t:= aRandomShape.(type) cast aRandomShape to the type
	case triangle:
		fmt.Printf("%T\n", t)
		fmt.Println("I got back a triangle with an area equal to ", getArea(t))
	case circle:
		fmt.Println("I got back a circle with an area equal to ", t.Area())
	default:
		fmt.Println("I don't recognize what I got back!")
	}

	var (
		x = 3
		y = 4.5
	)
	mySlice := []interface{}{x, y}
	z := mySlice[1]
	fmt.Printf("%T\n", z)
	var xx float64 = 4
	switch zz := z.(type) {
	case float64:
		fmt.Println(xx * zz) // cast to float64 type to allow for further processing
	default:
		fmt.Println("nothing")
	}
}
