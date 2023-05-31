package main

import "fmt"

type GitRemoteRepository struct {
	ProviderName string
}

func methods(methodsString string) {
	gh := GitRemoteRepository{
		ProviderName: methodsString,
	}

	WriteSpace()
	fmt.Println("Original value:", methodsString)

	var newValue string

	WriteSpace()
	fmt.Print("Provide new value: ")
	_, err := fmt.Scan(&newValue)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
}
