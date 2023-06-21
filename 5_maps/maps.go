/*
In this example, we make use of maps. We use maps (unlike arrays and slices) when we want a key
value pair without any order to our values
*/

package main

import "fmt"

func main() {
	// Creating an empty map
	students := make(map[string]int)

	// Adding values to the map
	students["Alice"] = 21
	students["Bob"] = 19
	students["Charlie"] = 20

	// Accessing values from the map
	fmt.Println("Alice's age:", students["Alice"])
	fmt.Println("Bob's age:", students["Bob"])

	// Modifying a value in the map
	students["Charlie"] = 21

	// Deleting a key-value pair from the map
	delete(students, "Bob")

	// Checking if a key exists in the map
	age, exists := students["Bob"]
	if exists {
		fmt.Println("Bob's age:", age)
	} else {
		fmt.Println("Bob is not in the map")
	}

	// Iterating over the map
	for name, age := range students {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}
