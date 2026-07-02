package main


import "fmt"


type ByteSize []byte


func (slice *ByteSize) Append(data []byte) {
	fmt.Printf("Before slice value in Append : %s\n", *slice)
	*slice = append(*slice, data...)
	fmt.Printf("After slice value in Append : %s\n ", *slice)
}


func main() {
	var b ByteSize = []byte("Hello ")
	fmt.Printf("Value of before calling Append : %s\n ", b)
	b.Append([]byte("World"))
	fmt.Printf("Value of b after calling Append : %s\n ", b)
}

