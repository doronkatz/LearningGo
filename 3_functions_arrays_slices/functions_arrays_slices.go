/*
In this exercise file, we will learn about both functions and making use of arrays and slices.
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	min := 0
	max := 100

	randomValue := getRandomNumber(min, max)
	fmt.Printf("Our random number (range from %v to %v, is %v\n", min, max, randomValue)

	// An array is an indexed collection of a certain size with values of the same type, declared as:
	// Source: https://gosamples.dev/capacity-and-length/
	randomArrayIntegers := [...]int{getRandomNumber(min, max), getRandomNumber(min, max), getRandomNumber(min, max), getRandomNumber(min, max)}
	fmt.Printf("randomIntegers: length: %d, capacity: %d, data: %v\n", len(randomArrayIntegers), cap(randomArrayIntegers), randomArrayIntegers)

	// A slice is not an array. It describes a section of the underlying array stored under the ptr pointer.
	randomSliceIntegers := []int{getRandomNumber(min, max), getRandomNumber(min, max), getRandomNumber(min, max), getRandomNumber(min, max)}
	fmt.Printf("randomIntegers: length: %d, capacity: %d, data: %v\n", len(randomSliceIntegers), cap(randomSliceIntegers), randomSliceIntegers)
	//set first value to zero
	randomSliceIntegers[0] = 0
	randomArrayIntegers[0] = 0

	fmt.Printf("randomSliceIntegers[0] %v\n", randomSliceIntegers[0])
	fmt.Printf("randomSliceIntegers[0] %v\n", randomSliceIntegers[0])

}

func getRandomNumber(min int, max int) int {
	return min + rand.Intn(max-min)
}
