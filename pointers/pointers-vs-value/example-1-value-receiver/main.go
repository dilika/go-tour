package main 


import "fmt"


type ByteSize []byte 

func (slice ByteSize) Append(data []byte) ByteSize {

	fmt.Printf("Slice value  in Append before %s\n ", slice)
	slice = append(slice, data...)
	fmt.Printf("Slice value in Append after %s\n ", slice)

	return slice
} 


func main() {
	var b ByteSize = []byte("Hello")
	fmt.Printf("Before Append Call b: %s\n", b)

	b = b.Append( []byte(" World"))
	fmt.Printf("After Append Call b: %s\n", b)

}
