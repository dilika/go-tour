package main

import "fmt"


func main(){

	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
 for i := range friends {
	 fmt.Printf("Value[%s]\t Adresse[%p]\t IndexAdress[%p]\n", friends[i], &i, &friends[i])
 }

}
