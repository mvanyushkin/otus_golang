package textanalyzer

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestWhenInputStringIsEmpty(t *testing.T) {
	result := GetMostUsedWords("")
	assert.Equal(t, 0, len(result))
}

func TestWhenSimpleRepeatableSequencePassed(t *testing.T) {
	result := GetMostUsedWords("cat cat cat cat catcat cat cat")
	assert.Equal(t, 2, len(result))
	assert.EqualValues(t, []string{"catcat", "cat"}, result)
}

func TestWhenLongWordSequencePassed(t *testing.T) {
	result := GetMostUsedWords(GetPseudoRandomString())
	assert.Equal(t, 10, len(result))
	assert.EqualValues(t, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, result)
}

func GetPseudoRandomString() string {
	var sb strings.Builder
	for index, v := range etalonWordsTable {
		for counter := 0; counter < index+1; counter++ {
			sb.WriteString(v)
			sb.WriteString(" ")
		}

		sb.WriteString(" - , : ; ")
	}

	return sb.String()
}

var etalonWordsTable = []string{
	"a",
	"b",
	"c",
	"d",
	"e",
	"f",
	"g",
	"h",
	"i",
	"j",
	"k",
	"l",
	"m",
	"o",
	"p",
	"q",
	"r",
	"s",
	"t",
	"u",
	"v",
}
