package main

import "fmt"

func main() {
    numbers1 := []int{1,2,3,4}
    fmt.Printf("The sum of %v is %d:\n", numbers1,foo(numbers1...))

    numbers2 := []int{5,6,7,8}
    fmt.Printf("The sum of %v is %d:\n", numbers2,bar(numbers2))
}

func foo (numbers ...int) int {
    sum := 0
    for _, number := range numbers {
       sum += number
    }
    return sum
}

func bar (numbers []int) int {
    sum := 0
    for _, number := range numbers {
       sum += number
    }
    return sum
}
