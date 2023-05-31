package main

import (
	"fmt"
)

func functions(functionsString string) {
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
