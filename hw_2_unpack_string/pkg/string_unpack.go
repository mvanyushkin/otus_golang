package unpackstring

import (
	"errors"
	"strconv"
	"unicode"
)

type StringUnpacker struct {
	lastSymbolRune            rune
	hasStartedEscapedSequence bool
	rawRepeatCount            string
	output                    string
}

func New() StringUnpacker {
	return StringUnpacker{}
}

func (_this *StringUnpacker) Do(input string) (string, error) {
	_, err := strconv.Atoi(input)
	if err == nil {
		return "", errors.New("wrong value, that is sucks")
	}

	for index, v := range input {
		// It covers the first iteration
		if _this.lastSymbolRune == 0 {
			_this.lastSymbolRune = v
			continue
		}

		if string(v) == "\\" {
			if _this.hasStartedEscapedSequence {
				_this.flushLastSymbol()
				_this.hasStartedEscapedSequence = false
				_this.lastSymbolRune = v
				continue
			} else {
				_this.hasStartedEscapedSequence = true
				continue
			}
		}

		if unicode.IsNumber(v) && !_this.hasStartedEscapedSequence {
			_this.rawRepeatCount += string(v)
			if index == len(input)-1 {
				_this.repeatLastSymbol()
			}
		}

		if _this.hasStartedEscapedSequence || !unicode.IsNumber(v) {
			if _this.rawRepeatCount == "" {
				_this.flushLastSymbol()
			} else {
				_this.repeatLastSymbol()
				_this.rawRepeatCount = ""
			}

			if index == len(input)-1 {
				_this.output += string(v)
				continue
			}

			_this.lastSymbolRune = v
			_this.hasStartedEscapedSequence = false
		}
	}

	result := _this.output
	_this.clearState()
	return result, nil
}

func (_this *StringUnpacker) repeatLastSymbol() {
	repeatString, _ := strconv.Atoi(_this.rawRepeatCount)
	for repeatIndex := 0; repeatIndex < repeatString; repeatIndex++ {
		_this.output += string(_this.lastSymbolRune)
	}
}

func (_this *StringUnpacker) flushLastSymbol() {
	_this.output += string(_this.lastSymbolRune)
}

func (_this *StringUnpacker) clearState() {
	_this.lastSymbolRune = 0
	_this.hasStartedEscapedSequence = false
	_this.rawRepeatCount = ""
	_this.output = ""
}
