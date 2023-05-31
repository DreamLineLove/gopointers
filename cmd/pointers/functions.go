package main

import (
	"fmt"
)

func functions(functionsString string) {
	// var functionsString string
	// flag.StringVar(&functionsString, "functionsString", "foo", "The value that will used in the basic() function.")
	// flag.Parse()
	WriteSpace()
	fmt.Println("Original value:", functionsString)

	var newValue string

	WriteSpace()
	fmt.Print("Provide new value: ")
	_, err := fmt.Scan(&newValue)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	changeValue(&functionsString, newValue)

	WriteSpace()
	fmt.Println("Value changed using function:", functionsString)

	WriteSpace()
	fmt.Println("Value accessed using *:", functionsString)
}

func changeValue(fs *string, newValue string) {
	*fs = newValue
}
