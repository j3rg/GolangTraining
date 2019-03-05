package main

import "fmt"

func main() {
    x := [][]string{ []string{"j3rg","testing"}, []string{"Jorgen","Ordonez"} }
    fmt.Println(x)

    xs1 := []string{"James", "Bond", "Shanken, not stirred"}
    xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
    fmt.Println(xs1)
    fmt.Println(xs2)

    xxs := [][]string{xs1,xs2}
    fmt.Println(xxs)

    for _, xs := range xxs {
        fmt.Println(xs)
    }
}
