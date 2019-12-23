package main

import (
	"fmt"
	"github.com/mvanyushkin/otus_golang/hw_3_most_frequent_words/textanalyzer"
)

func main() {
	res := textanalyzer.GetMostUsedWords("cat cat dog dog")
	for s := range res {
		fmt.Printf("%v \n", s)
	}
}
