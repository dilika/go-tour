package main

import "fmt"


func main(){

	var five [5]int

	four := [4]int{12, 24, 2, 5}

	// five = four // arrays/example2/main.go:12:9: cannot use four (variable of type [4]int) as [5]int value in assignment

	fmt.Println(five)
	fmt.Println(four)

}
