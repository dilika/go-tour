package main


import (
	"fmt"
	"encoding/json"
	"errors"

)

type user struct {
	ID int
	Name string
}

type updateStats struct {
	Modified int
	Duration float64
	Success bool
	Message string
}

func main(){
	u := user {
		ID: 123,
		Name: "Sally",
	}

	if _, err := updateUser(&u); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Updated user record for ID", u.ID)
}


func updateUser(u *user) (*updateStats, error){

	response := `{"Modified": 1, "Duration": 0.005, "Success": true, "Message": "modified wiht Success" }`

	var us updateStats
	if err := json.Unmarshal([]byte(response), &us); err != nil {
		return nil, err
	}

	if	us.Success != true {
		return nil, errors.New(us.Message)
	}

	return &us, nil
}
