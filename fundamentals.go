// In this module, you will learn about the basics of Go
package main

import "fmt"

func main() {
	// Declaring but not setting variables
	var aStringVariable string
	var anIntVariable int

	// Setting the variables
	aStringVariable = "hello"
	anIntVariable = 0

	for i := 0; i < 50; i++ {
		anIntVariable++
	}

	// Declaring and setting shorthand
	aFloatVariable := 2.0

	// printing to screen using fmt.Printf function.
	// https://pkg.go.dev/fmt#Printf
	fmt.Printf("aStringVariable = %v\n", aStringVariable)
	// Using ifs and for loops
	if aFloatVariable > 3.0 {
		fmt.Printf("aFloatVariable = %v\n", aFloatVariable)
	}

	fmt.Printf("anIntvariable = %v type= %T\n", anIntVariable, anIntVariable)

}
