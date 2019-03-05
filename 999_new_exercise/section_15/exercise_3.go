package main

import "fmt"

func main() {
    defer foo()

    fmt.Println("This line will print before the second line")
}

func foo() {
    fmt.Println("Second line")
}
