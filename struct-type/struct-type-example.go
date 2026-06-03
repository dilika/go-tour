package main

import "fmt"

type Example struct {
	flag    bool
	counter int16
	pi      float32
}

func main(){

	var  e1 Example

   fmt.Printf("%+v\n", e1)

	 e2 := Example {
		 flag :  true,
		 counter:  200,
		 pi     :  3.14,
		
	 }



	 fmt.Println("flag: ", e2.flag)
	 fmt.Println("Counter: ", e2.Counter)
	 fmt.Println("Pi: ", e2.pi)
 }

