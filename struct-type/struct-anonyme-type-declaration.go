package main

import "fmt"


func main(){

	var e struct {
		flag bool
		counter int16
		pi  float32
	}

  e1 := struct{

	
	
		flag bool
		counter int16
		pi float32
	} {
		flag:  false,
		counter: 500,
		pi: 3.14589030,
	} 

	 fmt.Println("flag: ", e1.flag)
	 fmt.Println("Counter: ", e1.counter)
	 fmt.Println("Pi: ", e1.pi)

	 fmt.Println("===========================")

	fmt.Printf("%+v\n", e)
}
