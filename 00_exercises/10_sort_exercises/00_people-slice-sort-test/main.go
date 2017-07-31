package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int      { return len(p) }
func (p people) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p people) Less(i, j int) bool { return p[i] < p[j] }

//func (p people) Less(i, j int) bool {
//str1 := []rune(strings.ToUpper(p[i]))
//str2 := []rune(strings.ToUpper(p[j]))

// Initialize lesserWord with str1
// Hence assuem str1 is lesserWord
//lesserWord := len(str1)

// Check if previous assumption was correct
//if len(str1) > len(str2) {
//		lesserWord = len(str2)
//}

// We need to know the lesser word so not to cause an index out of bound error
// e.g. word1 = Baton and word2 = Bat would cause an error if we looped based
// on the length of str1 as its of length 5 and str2 is of length 3
// an index out of bount error would occur on str2[3]
//for i := 0; i < lesserWord; i++ {
//if str1[i] != str2[i] {
//return str1[i] < str2[i]
//}
//}

//return true
//return strings.ToUpper(p[i]) < strings.ToUpper(p[j])
//}

func main() {
	//studyGroup := people{"Jenny", "Zac", "Joe", "Al"}
	studyGroup := people{"昨夜", "zeno", "Zeno", "john", "John", "al", "Al", "jenny", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}
