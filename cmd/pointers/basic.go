package main

import (
	"fmt"
)

func basic(basicString string) {
	// var basicString string
	// flag.StringVar(&basicString, "basicString", "foo", "The value that will used in the basic() function.")
	// flag.Parse()

	WriteSpace()
	fmt.Println("Original string:", basicString)

	var pointerToBasicString *string = &basicString

	WriteSpace()
	fmt.Print("Provide new value: ")
	_, err := fmt.Scan(pointerToBasicString)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	WriteSpace()
	fmt.Println("Value changed using pointer:", basicString)

	WriteSpace()
	fmt.Println("Value accessed using *:", *pointerToBasicString)
}
