package main

import "fmt"

func greatest(n ...int) int {
	var x int
	for _, y := range n {
		if y > x {
			x = y
		}
	}
	return x
}

func main() {
	x := []int{5, 7, 13, 4}
	fmt.Println("The greatest number of ", x, " is ", greatest(x...))
}
