package main

import "fmt"

func test(x func()) {
    x()
}

func main() {
    test(func(){
        fmt.Println("callback")
    })
}
