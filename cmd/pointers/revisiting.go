package main

import "fmt"

func revisiting() {
	var int1 int32
	int1 = 55
	var ptr1 *int32
	ptr1 = &int1
	fmt.Println("original value: ", int1)

	// Note that both work!
	changePtrByDereferencing(ptr1, 66)
	fmt.Println("new value 1: ", int1)
	changePtrByDereferencing(&int1, 99)
	fmt.Println("new value 2: ", int1)
	passingValues(int1, 100)
	fmt.Println("new value 3 (expect 99): ", int1)
}

func changePtrByDereferencing(intPtr *int32, newValue int32) {
	*intPtr = newValue
}

func passingValues(integer int32, newValue int32) {
	fmt.Println("original value when passed: ", integer)

	// this should only effect the function's local (parameter) copy of the value passed in
	integer = newValue
	fmt.Println("new value: ", integer)
}
