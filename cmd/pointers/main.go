package main

import "flag"

func main() {
	var basicsString string
	flag.StringVar(&basicsString, "basicsString", "foo", "The value that will be used in the basic() function.")
	var functionsString string
	flag.StringVar(&functionsString, "functionsString", "foo", "The value that will be used in the functions() function.")
	var methodsString string
	flag.StringVar(&methodsString, "methodsString", "foo", "The value that will be used in the methods() function.")
	flag.Parse()

	WriteTitle("With variables")
	basics(basicsString)

	NewLine()
	WriteTitle("With functions")
	functions(functionsString)

	NewLine()
	WriteTitle("With methods")
	methods(methodsString)

	NewLine()
	WriteTitle("Handling nil pointers")
	handlingNil()

	NewLine()
	WriteTitle("Revisiting pointers in Golang")
}
