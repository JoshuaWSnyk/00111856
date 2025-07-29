package main

import (
	"fmt"
	"rsc.io/quote" // Import the external dependency
)

func main() {
	// Use a function from the imported package
	fmt.Println("Here's a Go proverb:")
	fmt.Println(quote.Go()) // Prints a Go-related proverb

	fmt.Println("\nAnd here's another quote:")
	fmt.Println(quote.Hello()) // Prints "Hello, world."

	fmt.Println("\nAnd one about glass:")
	fmt.Println(quote.Glass()) // Prints a quote about glass
}