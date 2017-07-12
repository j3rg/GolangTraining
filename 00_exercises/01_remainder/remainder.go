package main

import "fmt"

func main() {
	var x int
	var y int

	fmt.Println("Enter number 1:")
	fmt.Scan(&x)

	fmt.Println("Enter number 2:")
	fmt.Scan(&y)

	result := x % y

	fmt.Printf("The remainder of %d and %d is %d\n", x, y, result)

}
