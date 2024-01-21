// interfaces/type-switch/begin/main.go
package main

import "fmt"

// define whatAmI which takes in an argument of any type and returns inforamtion about the underlying value's type
func whatAmI(a any) string {
	switch a.(type) {
	case int:
		return fmt.Sprintf("%v is an int", a)
	case string:
		return fmt.Sprintf("%v is a string", a)
	default:
		return fmt.Sprintf("%v is not supported (%T)", a, a)
	}
}

func main() {
	// invoke whatAmI function
	fmt.Println(whatAmI(1))
	fmt.Println(whatAmI("hello"))
	fmt.Println(whatAmI(true))
}
