package main


import "fmt"

func main() {

	// Create a slice with a length of 5 elements.
  fruits := make([]string, 5)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "coconuts"
	fruits[3] = "PineApple"
	fruits[4] = "Watermelon"

	// You can't access an index of a slice beyond its length.
  // fruits[5] = "Runtime errors"


	// Error: panic: runtime error: index out of range
	fmt.Println(fruits)
	fmt.Println(cap(fruits), len(fruits))
}
