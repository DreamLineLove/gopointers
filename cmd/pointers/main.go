package main

import "flag"

func main() {
	var basicString string
	flag.StringVar(&basicString, "basicString", "foo", "The value that will used in the basic() function.")
	var functionsString string
	flag.StringVar(&functionsString, "functionsString", "foo", "The value that will used in the basic() function.")
	flag.Parse()

	WriteTitle("With variables")
	basic(basicString)
	NewLine()
	WriteTitle("With functions")
	functions(functionsString)
}
