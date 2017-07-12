package main

import "fmt"

func zero(x int) {
	fmt.Printf("This is the memory address of variable x in func zero: %p\n", &x) // address in func zero
	fmt.Println(&x)                                                               //address in func zero
	x = 0
}

func main() {
	x := 5
	fmt.Printf("This is the memory address of variable x in func main: %p\n", &x) // address is main
	fmt.Println(&x)                                                               // address in main
	zero(x)
	fmt.Println(x) // x is still 5
}
