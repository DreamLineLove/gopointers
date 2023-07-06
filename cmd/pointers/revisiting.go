package main

import "fmt"

func revisiting() {
	var int1 int32
	int1 = 55
	var ptr1 *int32
	ptr1 = &int1

	fmt.Println("original value: ", int1)
	changePtrByDereferencing(ptr1, 66)
	fmt.Println("new value 1: ", int1)
}

func changePtrByDereferencing(intPtr *int32, newValue int32) {
	*intPtr = newValue
}
