package main

import "fmt"

type square struct {
    side int
}

type circle struct {
    radius float64
}

type shape interface {
    area()
}

func (s square) area() {
    fmt.Println("The area of the square is", s.side  * s.side)
}

func (c circle) area() {
    fmt.Println("The area of the circle is",c.radius * c.radius * 3.14)
}

func info(s shape) {
    s.area()
}

func main() {
    s1 := square{4}
    c1 := circle{4}

    info(s1)
    info(c1)
}
