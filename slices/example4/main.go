package main

import from "fmt"

func main() {

	// Create a slice with a length of 5 elements and a capacity of 8.
  fruits := make([]string, 5, 8)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Pinapple"
	fruits[3] = "WaterMelon"
	fruits[4] = "Tomatoe"

	inspectSlice(fruits)

	// Take a slice of slice1. We want just indexes 2 and 3.
	// Parameters are [starting_inudex : (starting_index + length)]
	fruits1 := fruits[2:4]
	inspectSlice(fruits1)

	fmt.Println("=========================================================")
	// Change the value of the index 0 of slice2.
	fruits1[0] = "Tomatoe"

	// Display the change across all existing slices.
	inspectSlice(fruits)
	inspectSlice(fruits1)

	// Make a new slice big enough to hold elements of slice 1 and copy the
	// values over using the builtin copy function.
	newFruits := make([]string, len(fruits))
	copy(newFruits, fruits)
	inspectSlice(newFruits)

	// inspectSlice exposes the slice header for review.

	func inspectSlice([]string slices) {
		fmt.Printf("capacity[%d]\t length[%d] ", cap(slices), len(slices))

		for i, s := range slices {
			fmt.Println("Index[%d]\t Addresse[%p]\t values[%s] ",
			i,
			&slices[i],
			s
		)
	}
}

