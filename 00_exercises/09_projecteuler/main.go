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
