package main

import (
	"flag"
	"fmt"
)

var functionsString string

func functions() {
	flag.StringVar(&functionsString, "functionsString", "foo", "The value that will used in the basic() function.")
	flag.Parse()
	WriteSpace()
	fmt.Println("Original value:", functionsString)
}
