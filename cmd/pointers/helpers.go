package main

import (
	"fmt"
	"strings"
)

func NewLine() {
	fmt.Println()
}

func WriteTitle(title string) {
	fmt.Println("\t", strings.ToUpper(title))
}

func WriteSpace() {
	fmt.Print(" ")
}
