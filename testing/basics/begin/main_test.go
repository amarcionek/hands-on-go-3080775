// testing/basics/begin/main_test.go
package main

import (
	"fmt"
	"testing"
)

// write a test for sum
func TestSum(t *testing.T) {
	actual := sum(1, 2, 3)
	expected := 6
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

// write a TestMain for setup and teardown
func TestMain(m *testing.M) {
	fmt.Println("**SETUP**")
	m.Run()
	fmt.Println("**TEARDOWN**")
}
