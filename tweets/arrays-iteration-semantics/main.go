package main

import "fmt"

// arrays/iteration semantics: does your loop SEE mutations
// you make to the array DURING iteration? Depends on which
// `for range` form you pick. This is the arrays keystone.

func main() {
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	// Form 1 — pointer semantics: iterate the ORIGINAL array.
	// Mutation is visible.
	fmt.Print("pointer (for i := range):  ")
	for i := range friends {
		friends[1] = "Jack"
		if i == 1 {
			fmt.Printf("friends[1] = %s\n", friends[1])
		}
	}

	// Form 2 — value semantics: `for range` COPIES the whole array
	// before iterating. v comes from the copy, made BEFORE the mutation.
	// Mutation is NOT visible in v.
	friends = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Print("value    (for i,v:=range): ")
	for i, v := range friends {
		friends[1] = "Jack"
		if i == 1 {
			fmt.Printf("v = %s\n", v)
		}
	}
}
