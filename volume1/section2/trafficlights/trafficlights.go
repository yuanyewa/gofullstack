// An example of declaring a constant grouping with an iota enumerator.
package main

import (
	"fmt"
)

type Light int
const (
	_ Light = iota
	TrafficLightStateRedLight
	TrafficLightStateGreenLight
	TrafficLightStateYellowLight
)

func (l Light) String() string {
	strings := []string {"Red", "Green", "Yellow"}
	return strings[l-1]
}
func main() {

	fmt.Println("Red Light State Code: ", TrafficLightStateRedLight)
	fmt.Println("Green Light State Code: ", TrafficLightStateGreenLight)
	fmt.Println("Yellow Light State Code: ", TrafficLightStateYellowLight)
}
