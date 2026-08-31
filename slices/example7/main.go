package main

import "fmt"

// user is a struct type that declares user information.
type user struct {
	id   int
	name string
}

func main() {

	// Declare and initialize a value of type user.
	u1 := user{
		id:   1,
		name: "Salvor Hardin",
	}

	// Declare and initialize a value of type user.
	u2 := user{
		id:   2,
		name: "Gaal Dornick",
	}

	// Display both user values.
	display(u1, u2)

	// Create a slice of user values.
	u3 := []user{
		{4, "Johran Sutt"},
		{5, "Lewis Pirenne"},
	}

	// Display all the user values from the slice.
	display(u3...)

	change(u3...)

	fmt.Println("******************************< U3 display >****************************")
	for _, u := range u3 {
		fmt.Printf("%+v\n", u)
	}
}

// display can accept and display multiple values of user types.
func display(users ...user) {
	fmt.Println("******************************< User Display >***************************")
	for _, u := range users {
		fmt.Printf("User id[%d] and value[ %v ] \n", u.id, u)
	}
}

// change shows how the backing array is shared.
func change(users ...user) {
	users[1] = user{6, "Hobert Manlio"}
}
