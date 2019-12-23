package textanalyzer

import (
	"sort"
	"strings"
)
import f "github.com/thoas/go-funk"

const DefaultWordsLimit = 10

var excludedSymbols = []string{"-", ",", ";", ":", "."}

type AnalyzerDict = map[string]uint64

type EntryCount struct {
	value string
	count uint64
}

func GetMostUsedWords(value string, limit ...int) []string {
	actualLimit := DefaultWordsLimit
	if len(limit) > 0 {
		actualLimit = limit[0]
	}

	for _, e := range excludedSymbols {
		value = strings.Replace(value, e, " ", -1)
	}

	words := strings.Split(value, " ")[:]
	var measurableWords interface{}
	measurableWords = f.Filter(words, func(x string) bool {
		return x != ""
	})

	frequentMap := make(AnalyzerDict)
	for _, v := range measurableWords.([]string) {
		frequentMap[v]++
	}

	entries := make([]EntryCount, len(frequentMap))
	sliceIndex := 0
	for v, c := range frequentMap {
		entry := EntryCount{
			value: v,
			count: c,
		}
		entries[sliceIndex] = entry
		sliceIndex++
	}

	sort.Slice(entries, func(left, right int) bool {
		return entries[left].count < entries[right].count
	})

	var mostUsedWords interface{}
	mostUsedWords = f.Map(entries, func(x EntryCount) string {
		return x.value
	})

	mostUsedWordsSlice := mostUsedWords.([]string)
	if len(mostUsedWordsSlice) <= actualLimit {
		return mostUsedWordsSlice
	}

	return mostUsedWordsSlice[0:actualLimit]
}
