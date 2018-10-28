package mypackage

import (
	"fmt"
)

var XX int

type MyS struct {
	Name    string
	Age     int
	Address []string
}

func (m MyS) Update(name string, age int, address []string) *MyS {
	m = MyS{Name: name, Age: age, Address: address}
	return &m
}

func init() {
	XX = 3
	fmt.Println("myPackage is called")
}
