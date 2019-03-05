package main

import "fmt"

type person struct {
    first string
    last string
    age int
}

func (p person) speak() {
    fmt.Println("Hi, my name is ", p.first, p.last)
}

func main() {
    x := person{"Jorgen", "Ordonez", 30}
    x.speak()
}
