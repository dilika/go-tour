package main 


import (
	"fmt"
)

type ByteSlice []byte

func (slice *ByteSlice) Write(data []byte) (n int , err error) {
	*slice = append(*slice, data...)
	return len(data), nil
}


func main(){
	var b ByteSlice
	fmt.Fprintf(&b, "this hour has %d days ", 7)
	fmt.Printf("Final result : %s ", b)
}
