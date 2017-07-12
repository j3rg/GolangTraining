package main

import "fmt"

// NOTE: The AND operator hasn't been thought to us
// as yet hence there is no use of it
func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("Fizz")
			if i%5 == 0 {
				fmt.Printf("Buzz")
			}
		} else {
			if i%5 == 0 {
				fmt.Printf("Buzz")
			} else {
				fmt.Printf("%d", i)
			}
		}
		fmt.Printf("\n")
	}
}
