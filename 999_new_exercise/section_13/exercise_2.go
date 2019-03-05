package main

import "fmt"

type person struct {
    firstname string
    lastname string
    favorite_flavors []string
}

func main() {

    x := person{
        firstname: "Jorgen",
        lastname: "Ordonez",
        favorite_flavors: []string{"Vanilla", "Cheese Cake"},
    }

    fmt.Println("First Name: ", x.firstname)
    fmt.Println("Last Name: ", x.lastname)
    fmt.Println("Favorite Ice Cream Flavors: ")
    for _, v := range x.favorite_flavors {
        fmt.Printf("\t%v\n",v)
    }

    y := person{
        firstname: "Alexi",
        lastname: "Ordonez2",
        favorite_flavors: []string{"Oreo Cookie", "Chocolate"},
    }

    fmt.Println("First Name: ", y.firstname)
    fmt.Println("Last Name: ", y.lastname)
    fmt.Println("Favorite Ice Cream Flavors: ")
    for _, v := range y.favorite_flavors {
        fmt.Printf("\t%v\n",v)
    }

    m := map[string]person{
        x.lastname: x,
        y.lastname: y,
    }

    fmt.Println(m)

}
