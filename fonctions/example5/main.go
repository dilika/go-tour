package main

import (
	"fmt"
	"runtime"
)

func main(){
	if err := testPanic(); err != nil {
		fmt.Println("Error ", err)
	}
}


func testPanic() (err error){

	defer catchPanic(&err)

	fmt.Println(" Start test ")

	err = mimicError("1")

	var p *int 

	*p = 10

	fmt.Println("End TEst")
	return err
}


func catchPanic(err *error) {

	if r := recover(); r != nil {
		fmt.Println("Panic DEFERED")

		buff := make([]byte, 10000)
		runtime.Stack(buff, false)
		fmt.Println("Stack trace : ", string(buff))

		if err != nil {
			*err = fmt.Errorf("%v", r)
		}

	}
}


func mimicError(key string) error {
	return fmt.Errorf("Mimic error : %s ", key)
}
