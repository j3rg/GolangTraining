package main

import "fmt"

func half(n int) (int, bool) {
	if n%2 == 0 {
		return n / 2, true
	}

	return n / 2, false
}

func main() {

	funcExp := func(n int) (int, bool) {
		if n%2 == 0 {
			return n / 2, true
		}

		return n / 2, false
	}
	x, y := funcExp(3)
	fmt.Println("half(3) returns ", x, " and ", y)
	x, y = funcExp(4)
	fmt.Println("half(4) returns ", x, " and ", y)
}
