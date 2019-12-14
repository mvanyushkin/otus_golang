package unpackstring

import (
	"errors"
	"strconv"
	"unicode"
)

func DoUnpackString(input string) (string, error) {
	var repeatableRune rune
	var rawRepeatCount string
	var output = ""
	_, err := strconv.Atoi(input)
	if err == nil {
		return "", errors.New("wrong value, that is sucks")
	}

	for index, v := range input {
		// It covers the first iteration
		if repeatableRune == 0 {
			repeatableRune = v
			continue
		}

		if unicode.IsNumber(v) {
			rawRepeatCount += string(v)
		}

		if !unicode.IsNumber(v) {
			if rawRepeatCount == "" {
				output += string(repeatableRune)
			} else {
				repeatString, _ := strconv.Atoi(rawRepeatCount)
				for repeatIndex := 0; repeatIndex < repeatString; repeatIndex++ {
					output += string(repeatableRune)
				}

				rawRepeatCount = ""
			}

			if index == len(input)-1 {
				output += string(v)
				continue
			}

			repeatableRune = v
		}
	}

	return output, nil
}
