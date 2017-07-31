package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"zeno", "Zeno", "john", "John", "al", "Al", "jenny", "Jenny"}

	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
}
