package main

import (
	"fmt"
	"unsafe"
)

type ByteSlice []byte

func (slice ByteSlice) ValueMethode() {
	//slice = append(slice, []byte(" World")...)
	fmt.Println("ValueMethode called on ", string(slice))
}

func (slice *ByteSlice) PointerMethode() {
	*slice = append(*slice, []byte("!!!")...)
	fmt.Println("PointerMethode called on ", string(*slice))
}

func main() {
	var b ByteSlice = []byte("Hello")

	fmt.Println("============= Call on addressable variables ======================")
	b.ValueMethode()
	b.PointerMethode()

	fmt.Println("============= Call on litteral variables =========================")
	ByteSlice("Test").ValueMethode()



	fmt.Println("\n\n============= Litteral variables Test ============================")

	var litteral ByteSlice = []byte("Test")
	fmt.Printf("Addresse of litteral : %p\n", &litteral)

	fmt.Printf("Header Addresse of litteral : %p\n", unsafe.Pointer(&litteral))

	fmt.Printf("Addresse of data  : %p\n", unsafe.Pointer(&litteral[0]))

}
