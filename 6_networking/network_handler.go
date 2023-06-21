/* In this example, I make an HTTP GET call to a website, check for errors,
and then print the response body.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, error := http.Get("https://doronkatz.com")

	if error != nil {
		fmt.Printf("Error making GET request: %s\n", error)
		return //end program run
	}

	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	// Print the response body
	fmt.Println("Response body:")
	fmt.Println(string(body))
}
