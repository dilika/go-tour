// Declare a struct type to maintain information about a user. Declare a function
// that creates value of and returns pointers of this type and an error value. Call
// this function from main and display the value.
//
// Make a second call to your function but this time ignore the value and just test
// the error value.

package main

// Add imports.

import (
	"fmt"
)

// Declare a type named user.

type user struct {
	name string
	lastName string
	age int
}

// Declare a function that creates user type values and returns a pointer
// to that value and an error value of nil.

func createNewUser() (*user, error)/* (pointer return arg, error return arg) */ {

	newUser := user {
		name: "Salvor",
		lastName: "Harding",
		age : 200,
	}

	return &newUser, nil
	// Create a value of type user and return the proper values.
}

func main() {

	// Use the function to create a value of type user. Check
	userFormFunc, error := createNewUser()
	// the error being returned.
	if error != nil {
		fmt.Println("Erreur while creating : ", error)
		return
	}

	// Display the value that the pointer points to.

	fmt.Printf("The new User created is:  %+v", *userFormFunc)

	// Call the function again and just check the error.

	if _, error := createNewUser(); error != nil {
		fmt.Println("Erreur called just with check error with first value returned ", error)
		return
	}
}
