package main

import "fmt"

func main() {
    var x [5]int

    x[0] = 3
    x[1] = 4
    x[2] = 5
    x[3] = 6
    x[4] = 7

    for i, v := range x {
        fmt.Println(i,v)
    }

    fmt.Printf("%T\n",x)
}

