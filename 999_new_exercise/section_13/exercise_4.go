package main

import "fmt"

func main() {
    p := struct{
        fname string
        lname string
        hobbies []string
        siblings map[string]int
    }{
        fname: "Jorgen",
        lname: "Ordonez",
        hobbies: []string{
            "Hacking","Traveling","Cooking","Development",
        },
        siblings: map[string]int{
            "Borthers":2,
            "Sisters":0,
        },
    }

    fmt.Println(p)
}
