package main

import "fmt"

func sayHello() func() {
    return func() {
        fmt.Println("Hello")
    }
}
func main() {
    x := sayHello()
    x()
}
