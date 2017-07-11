package main

import "fmt"

const (
	_  = iota             //0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	fmt.Println("Byte measurements:")
	fmt.Println("binary\t\t\t\t\t\tdecimal\t\t\tprefix")
	fmt.Printf("%b\t\t\t\t\t", KB)
	fmt.Printf("%d\t\t\t", KB)
	fmt.Printf("Kilo\n")
	fmt.Printf("%b\t\t\t\t", MB)
	fmt.Printf("%d\t\t\t", MB)
	fmt.Printf("Mega\n")
	fmt.Printf("%b\t\t\t", GB)
	fmt.Printf("%d\t\t", GB)
	fmt.Printf("Giga\n")
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\t\t", TB)
	fmt.Printf("Tera\n")
}
