package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)

	m := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Printf("The type of m is %T\n", m)
	fmt.Printf("The address of m is %p\n", m)
	fmt.Println(m)

	fmt.Println("Creating IntSlice for m...")
	intSlice := sort.IntSlice(m)
	fmt.Printf("The type of intSlice is %T\n", intSlice)
	fmt.Printf("The address of intSlice is %p\n", intSlice)
	fmt.Printf("The value of &intSlice %p\n", &intSlice)
	fmt.Println(intSlice)

	fmt.Println("Reversing IntSlice ...")
	revIntSlice := sort.Reverse(intSlice)
	fmt.Printf("The type of revIntSlice is %T\n", revIntSlice)
	fmt.Printf("The address of revIntSlice is %p\n", revIntSlice)
	fmt.Printf("The value of &revIntSlice is %p\n", &revIntSlice)
	fmt.Println(revIntSlice)

	fmt.Println("Sorting revIntSlice...")
	sort.Sort(revIntSlice)
	fmt.Printf("The type of n is %T\n", m)
	fmt.Printf("The address of n is %p\n", m)
	fmt.Println(m)
	//j := &i
	//fmt.Printf("The type at address %p is %T with the address of %p", i, &i, j)
}
