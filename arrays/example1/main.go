package main

import "fmt"

func main(){

	var fruits [5]string

	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Pinapple"
	fruits[3] = "WaterMellon"
	fruits[4] = "Tomatoe"

	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	fmt.Printf("\n=========================== languages ========================\n")

	languages := [5]string{"Go", "Java", "JavaScript", "Python", "Rust"}

	for i := range languages {
		fmt.Println(i, languages[i])
	}
}
