package main

import (
	"fmt"
	"hw_3_most_frequent_words/textanalyzer"
)

type Person struct {
	id int
}

func main() {
	res := textanalyzer.GetMostUsedWords("cat cat dog dog")
	for s := range res {
		fmt.Printf("%v \n", s)
	}
}
