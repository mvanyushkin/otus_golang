package main

import (
	"fmt"
	us "hw_2_unpack_string/pkg"
)

type Foo struct {
	id int
}

func (f *Foo) Do() {
	f.id = 8
}

func main() {
	unpacker := us.New()
	printout(&unpacker, "a4bc2d5e")
	printout(&unpacker, "abcd")
	printout(&unpacker, "45")
	printout(&unpacker, "qwe\\4\\5")
	printout(&unpacker, "qwe\\45")
	printout(&unpacker, "qwe\\\\5")
	printout(&unpacker, "a10bc2d5e")
}

func printout(unpacker *us.StringUnpacker, testString string) {
	result, err := unpacker.Do(testString)
	if err != nil {
		fmt.Printf("Input %v has led to the exception: \n", err.Error())
	} else {
		fmt.Printf("Input: %v | Output: %v \n", testString, result)
	}
}
