package main

import (
	"fmt"
	"errors"
)


func divide(a int, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("Can't divide a number by 0")
	}

	return a/b, nil

}


func main(){
	numerator := 20
	denom := 5

	result, error := divide(numerator, denom)

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(result)
}
