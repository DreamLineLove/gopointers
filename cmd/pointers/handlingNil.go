package main

import "fmt"

func handlingNil() {
	var nilBooleanPointer *bool

	changeBooleanValue(nilBooleanPointer)

	var trueBool bool = true
	WriteSpace()
	fmt.Println("Original value:", trueBool)

	nilBooleanPointer = &trueBool
	changeBooleanValue(nilBooleanPointer)

	WriteSpace()
	fmt.Println("Changed value:", trueBool)

	NewLine()
	WriteSpace()
	fmt.Println("-\tRead the source code for handlingNil.go to learn more")
}

func changeBooleanValue(boolean *bool) {
	// if the Pointer boolean happens to be nil, the function will exit with a standard output
	if boolean == nil {
		WriteSpace()
		fmt.Println("ERROR occurred when trying to access the pointer")
		return
	}

	// if not it will perform the intended logic
	*boolean = !*boolean
}
