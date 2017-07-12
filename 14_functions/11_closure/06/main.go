package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	fmt.Printf("%T", callback)
	return xs
}

func main() {
	y := 5
	xs := filter([]int{1, 2, 3, 4, y}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs) // [2 3 4]
}
