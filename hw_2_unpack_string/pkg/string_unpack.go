package unpackstring

import (
	"errors"
	"strconv"
	"unicode"
)

func DoUnpackString(input string) (string, error) {
	var lastSymbolRune rune
	var hasStartedEscapedSequence = false
	var rawRepeatCount string
	var output = ""
	_, err := strconv.Atoi(input)
	if err == nil {
		return "", errors.New("wrong value, that is sucks")
	}

	for index, v := range input {
		// It covers the first iteration
		if lastSymbolRune == 0 {
			lastSymbolRune = v
			continue
		}

		if string(v) == "\\" {
			if hasStartedEscapedSequence {
				output += string(lastSymbolRune)
				hasStartedEscapedSequence = false
				lastSymbolRune = v
				continue
			} else {
				hasStartedEscapedSequence = true
				continue
			}
		}

		if unicode.IsNumber(v) && !hasStartedEscapedSequence {
			rawRepeatCount += string(v)
			if index == len(input)-1 {
				repeatString, _ := strconv.Atoi(rawRepeatCount)
				for repeatIndex := 0; repeatIndex < repeatString; repeatIndex++ {
					output += string(lastSymbolRune)
				}
			}
		}

		if hasStartedEscapedSequence || !unicode.IsNumber(v) {
			if rawRepeatCount == "" {
				output += string(lastSymbolRune)
			} else {

				repeatString, _ := strconv.Atoi(rawRepeatCount)
				for repeatIndex := 0; repeatIndex < repeatString; repeatIndex++ {
					output += string(lastSymbolRune)
				}

				rawRepeatCount = ""
			}

			if index == len(input)-1 {
				output += string(v)
				continue
			}

			lastSymbolRune = v
			hasStartedEscapedSequence = false
		}
	}

	return output, nil
}
