package main

import "fmt"

type user struct {
	likes int
}

func main() {

// Declare a slice of 3 users.
users := make([]user, 3)

// Share the user at index 1.
shareUser := &users[1]

// Add a like for the user that was shared.

shareUser.likes++
// Display the number of likes for all users.

fmt.Printf("\n======================< Before append >========================\n")
for i := range users {
	fmt.Printf("user [%d] has number of like [%d] \n", i, users[i].likes)
}

// Add a new user.
users = append(users, user{})

// Display the number of likes for all users.
fmt.Printf("\n======================< after append >========================\n")
for i := range users {
	fmt.Printf("user [%d] has number of likes [%d] \n", i, users[i].likes)
}


// Notice the last like has not been recorded.

}
