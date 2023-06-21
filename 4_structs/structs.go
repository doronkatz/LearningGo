/*
In this exercise file, we will learn about working with structs
*/

package main

import "fmt"

type Car struct {
	Color string
	VIN   string
	make  string
	model string
}

func main() {
	ford := Car{"Green", "123546708", "Ford", "Falcon"}
	fmt.Printf("Ford color: %v make: %v, model %v", ford.Color, ford.make, ford.model)
}
