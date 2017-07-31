package main

import (
	"fmt"
	"sort"
	"strings"
)

type people []string

func (p people) Len() int { return len(p) }

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p people) Less(i, j int) bool {
	str1 := []rune(strings.ToUpper(p[i]))
	str2 := []rune(strings.ToUpper(p[j]))

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return str1[i] < str2[i]
		}
	}

	return false
}

func main() {

	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

}
