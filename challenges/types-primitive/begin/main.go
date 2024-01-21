// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func main() {
	// create a local variable "val" and assign it an int value
	val := 42

	// print the value of the local variable "val"
	fmt.Printf("Local val: %v, Type %T\n", val, val)

	// print the value of the package-level variable "val"
	fmt.Printf("Package val: %v, Type %T\n", getVal(), getVal())

	// update the package-level variable "val" and print it again
	updateVal("newGlobal")
	fmt.Printf("Package val: %v, Type %T\n", getVal(), getVal())

	// print the pointer address of the local variable "val"
	var aofVal *int = &val
	fmt.Printf("Address of val: %v\n", aofVal)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 100
	fmt.Printf("Local val: %v, Type %T\n", val, val)
}

func getVal() string {
	return val
}

func updateVal(newVal string) string {
	val = newVal
	return val
}
