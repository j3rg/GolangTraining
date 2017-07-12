package main

import "fmt"

func half(n int) (int, bool) {
	if n%2 == 0 {
		return n / 2, true
	}

	return n / 2, false
}

func main() {
	x, y := half(3)
	fmt.Println("half(3) returns ", x, " and ", y)
	x, y = half(4)
	fmt.Println("half(4) returns ", x, " and ", y)
}
