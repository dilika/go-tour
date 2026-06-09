package main

import "fmt"

type gc struct {
	newFlag    bool
	counter int16
	pi      float32
}

func main(){

	var  e1 gc

   fmt.Printf("%+v\n", e1)

	 e2 := gc {
		 newFlag :  true,
		 counter:  200,
		 pi     :  3.14,
		
	 }


	 fmt.Println("Counter: ", e2.Counter)
	 fmt.Println("newFlag: ", e2.newFlag)
	 fmt.Println("Pi: ", pi)
 }

