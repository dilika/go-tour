package main

import "fmt"

func main() {

	// Create a slice with a length of 5 elements and a capacity of 8.
	fruits := make([]string, 5, 8)
	fruits[0] = "Apple"
	fruits[1] = "Tomato"
	fruits[2] = "Orange"
	fruits[3] = "Watermelon"
	fruits[4] = "PineApple"


// inspectSlice exposes the slice header for review.
	inspectSlice(fruits)
}


func inspectSlice(fruits []string) {
	fmt.Printf("length[%d] capacity[%d]\n ", len(fruits), cap(fruits))
	for i := range fruits {
		fmt.Printf("[%d] %p %s\n",
		i,
		&fruits[i],
		fruits[i])
	}
}
