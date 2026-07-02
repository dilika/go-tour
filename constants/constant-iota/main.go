package main

import "fmt"


func main(){

	const (
		A = iota
		B
		C
		D = iota
	)

	fmt.Println(A, B, C, D)
}
