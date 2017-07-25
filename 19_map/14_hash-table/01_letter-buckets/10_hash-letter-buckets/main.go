package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book moby dick
	// the original URL from the lesson does not work anymore
	//res, err := http.Get("http://www.gutenberg.org/files/2017/old/moby10b.txt")
	res, err := http.Get("http://www.gutenberg.org/files/2017/2017.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Create slice to hold counts
	buckets := make([]int, 200)
	// Loop over the words
	for scanner.Scan() {
		n := hashBucket(scanner.Text())
		buckets[n]++
	}
	fmt.Println(buckets[65:123])
	// fmt.Println("******************")
	// for i := 28; i <= 126; i++ {
	// 	fmt.Printf("%v - %c - %v\n", i, i, buckets[i])
	// }

	for i, count := range buckets[65:123] {
		// need as buckets[65:123] generates a new slice starting at index 0
		i += 65
		fmt.Println(i, " - ", string(i), " - ", count)
	}
}

func hashBucket(word string) int {
	return int(word[0])
}
