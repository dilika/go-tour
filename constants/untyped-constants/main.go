package main

import "fmt"

func main(){
	const (
		Answer = 42
		Pi 		= 3.14
		Name		= "Golang"
	)


	var x int = Answer
	var y  float64 = Answer
	var z		string	= Name

	fmt.Println(x, y, z)
}
