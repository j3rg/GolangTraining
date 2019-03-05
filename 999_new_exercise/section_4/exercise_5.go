package main

import "fmt"

type pizza int
var x pizza
var y int

func main() {
    fmt.Printf("%v\n",x)
    fmt.Printf("%T\n",x)
    x = 42
    fmt.Printf("%v\n",x)
    y = int(x)
    fmt.Println(y)
    fmt.Printf("%T\n",y)
}

