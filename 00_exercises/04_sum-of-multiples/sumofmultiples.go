package main

import "fmt"

func main() {
	var x int
	for i := 3; i < 1000; i++ {
		if i%3 == 0 {
			x += i
		} else if i%5 == 0 {
			x += i
		}
	}
	fmt.Printf("The sum of multiples of 3 and 5 below a 1000 is: %d\n", x)
}
