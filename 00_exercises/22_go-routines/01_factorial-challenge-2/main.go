package main

import "fmt"

func main() {
	c1 := channelSetter(4, 5)
	c2 := factorial(c1)
	for n := range c2 {
		fmt.Println(n)
	}
}

func channelSetter(n ...int) chan int {
	out := make(chan int)
	go func() {
		for i := range n {
			out <- n[i]
		}
		close(out)
	}()
	return out
}

func factorial(c chan int) chan int {
	out := make(chan int)
	go func() {
		for x := range c {
			out <- fact(x)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
