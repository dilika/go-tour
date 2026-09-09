package main

import "fmt"

func main() {
	// Using the value semantic form of the for range.
	friends := []string{"Hardin", "Lewis", "Manlio", "Pirenne", "Dornick", "Sutt"}
	for _, v := range friends {
		friends = friends[:2]
		fmt.Printf("\nHere is the friend : [ %v ] \t ", v)
	}

	friendss := []string{"Hardin", "Lewis", "Manlio", "Pirenne", "Dornick", "Sutt"}
	for i := range friendss {
		friendss = friendss[:2]
		fmt.Printf("\n Here is friend [%v]", friends[i])
	}

	// Using the pointer semantic form of the for range.
}
