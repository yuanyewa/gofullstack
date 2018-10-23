// An example of using Go's iota enumerator to model the year and city where the summer Olympics took place.
package main

import (
	"fmt"
	"strconv"
)

// iota increate by 1 each time; omit expression will adopt from previous one
type myConst int

const (
	x myConst = 10 + iota
	y         = 20 + iota
	_
	z
)

type myInt int

func (m myInt) String() string {
	return strconv.Itoa(int(m) * 10)
}

// type that implements a String() function can be printed accordingly
func (m myConst) String() string {
	switch m {
	case x:
		return "myX"
	case y:
		return "y"
	case z:
		return "z"
	default:
		return "nothing"
	}
}

func main() {
	fmt.Println(myInt(20))
	fmt.Println(x)
	const (
		LosAngeles = 1984 + (iota * 4)
		Seoul
		Barcelona
		Atlanta
		Sydney
		Athens
		Beijing
		London
		Rio
		Tokyo
	)

	fmt.Println("These cities hosted or will host the Summer Olympics in the provided year...")
	fmt.Printf("%-18s %-18s \n", "City", "Year")
	fmt.Printf("%-18s %-18v \n", "Los Angeles", LosAngeles)
	fmt.Printf("%-18s %-18v \n", "Atlanta", Atlanta)
	fmt.Printf("%-18s %-18v \n", "Tokyo", Tokyo)

}
