/*
In this exercise file, you will learn about using strings, working with sub-strings and more.
*/

package main

import "fmt"

func main() {
	brownFox := "The quick brown fox jumps over the lazy dog"

	//Using a string, we can get a substring using the [:] slicing notation
	fmt.Printf("The first five letters are '%v'\n", brownFox[:5])

	//To get length of the sentence len (https://pkg.go.dev/builtin#len)
	length := len(brownFox) - 1
	fmt.Printf("the length of the sentence is %v", length)
}
