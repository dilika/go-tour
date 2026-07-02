package main 


import "fmt"

func main(){

	const (
		PI = 3.14159
		TypedPI float64 = 3.14159
	)


	fmt.Println(PI, TypedPI)

	var rayon int = 10
	air := TypedPI * float64(rayon) * float64(rayon)

	fmt.Println(air)
}
