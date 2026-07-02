package main

// From Spec:
// a short variable declaration may redeclare variables provided they
// were originally declared earlier in the same block with the same
// type, and at least one of the non-blank variables is new.

// Sample program to show some of the mechanics behind the
// short variable declaration operator redeclares.

import "fmt"

type user struct {
	id int
	name string
	}


	func main() {


		var err1 error


		u, err1 := getUser()
		if err1 != nil {
			return
		}

		fmt.Println(u)

		u, err2 := getUser()
		if err2 != nil {
			return
		}

		fmt.Println(u)

	}



	func getUser() (*user, error) {
			return &user{123, "Betty"}, nil
	}
