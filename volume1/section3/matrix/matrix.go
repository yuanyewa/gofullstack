// An example of a multidimensional array, a 3x4 matrix.
package main

import (
	"fmt"
)

func main() {
	myMat := [3][]int{{1, 2}, {3, 4, 5}}
	myMap2 := make([][]int, 3, 4)
	// myMap2[0] = []int{1, 2}
	fmt.Println("capacity & len is: ", cap(myMap2), len(myMap2))
	fmt.Println(myMat)
	for i := 0; i < len(myMat); i++ {
		for j := 0; j < len(myMat[i]); j++ {
			fmt.Printf("%5d", myMat[i][j])
		}
		fmt.Printf("\n")
	}

	// Here we create a multi-dimensional array, a 3x4 matrix (3 rows, 4 columns)
	myMatrix := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

	// Let's print a value from a cell in each row of the matrix
	fmt.Println("Value at [0][0]: ", myMatrix[0][0])
	fmt.Println("Value at [1][1]: ", myMatrix[1][1])
	fmt.Println("Value at [2][2]: ", myMatrix[2][2])
	fmt.Println("\n")

	// First attempt at printing the matrix
	fmt.Printf("%+v\n", myMatrix)
	fmt.Println("\n")

	// A better looking matrix output to the screen
	printThreeByFourMatrix(myMatrix)
}

// A function to print the 3x4 matrix in a more pretty manner
func printThreeByFourMatrix(inputMatrix [3][4]int) {

	rowLength := len(inputMatrix)
	columnLength := len(inputMatrix[0])

	for i := 0; i < rowLength; i++ {
		for j := 0; j < columnLength; j++ {
			fmt.Printf("%5d ", inputMatrix[i][j])
		}
		fmt.Println()
	}

}
