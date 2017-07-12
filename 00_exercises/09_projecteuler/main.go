package main

import "fmt"

// A diagonal line can only be done using
// and even row and column creating a perfectly
// squared grid of numbers hence only one input
func numberSpiralDiagonals(n int) int {
	// Calculate number to start spiral
	startNum := n * n
	var sum int

	// Loop until starting number is one or zero
	// as this indicates the end of the spiral
	// 1 for odd numbers 0 for even numbers
	for startNum > 1 {
		// subtract one from number to calculate corner number
		n--
		// Add corner numbers and start corner of next
		for i := 0; i <= 3; i++ {
			sum += startNum
			startNum -= n
		}
		// recalculate for inner grid
		n--
	}
	// Return sum plus remaining value
	// remaing value can only be 1 for odd numbers or 0 for even numbers
	return sum + startNum
}

func main() {
	var x int
	fmt.Print("Please enter a single number: ")
	fmt.Scan(&x)
	y := numberSpiralDiagonals(x)
	fmt.Println("The spiral diagonal sum of ", x, " is ", y)
}

// Problem 28: Number spiral diagonals
// Reference: https://projecteuler.net/problem=28
//
/*
 * Starting with the number 1 and moving to the right in a clockwise direction
 * a 5 by 5 spiral is formed as follows:
 *
 * [21]  22  23  24  [25]
 *  20  [7]   8  [9]  10
 *  19   6   [1]  2   11
 *  18  [5]   4  [3]  12
 * [17] 16   15  14  [13]
 *
 * It can be verified that the sum of the numbers on the diagonals is 101.
 * What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral
 * formed in the same way?
 */
