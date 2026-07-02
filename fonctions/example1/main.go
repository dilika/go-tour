
package main

// named and struct types.
// Sample program to show how functions can return multiple values while using

import (
	"fmt"
	"encoding/json"
)


type user struct {
	ID int
	name string
}


func main(){

	user := struct {
		ID int
		name string
	} {
		ID : 123,
		name : "Sally",
	}

	u, err := retrieveUser(user.name)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", *u)

}


func retrieveUser(name string) (*user, error){

	r, err := getUser(name)

	if err != nil {
		return nil, err
	}

	var u user
	error := json.Unmarshal([]byte(r), &u)

	return &u, error
}



func getUser(name string) (string, error) {
	response := `{"id":123, "name":"sally"}`
	return response, nil
}
