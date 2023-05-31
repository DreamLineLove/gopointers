package main

import "flag"

func main() {
	var basicString string
	flag.StringVar(&basicString, "basicString", "foo", "The value that will be used in the basic() function.")
	var functionsString string
	flag.StringVar(&functionsString, "functionsString", "foo", "The value that will be used in the functions() function.")
	var methodsString string
	flag.StringVar(&methodsString, "methodsString", "foo", "The value that will be used in the methods() function.")
	flag.Parse()

	WriteTitle("With variables")
	basics(basicString)
	NewLine()
	WriteTitle("With functions")
	functions(functionsString)
	NewLine()
	WriteTitle("With methods")
	methods(methodsString)
}
