package main

import "fmt"

type pizza int
var x pizza

func main() {
    fmt.Printf("%v\n",x)
    fmt.Printf("%T\n",x)
    x = 42
    fmt.Printf("%v\n",x)
}

