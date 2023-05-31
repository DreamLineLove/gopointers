package main

import (
	"fmt"
)

func basics(basicString string) {
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
