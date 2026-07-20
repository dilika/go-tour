package main


import "fmt"

func main() {

	// Declare a nil slice of strings.
	var nilSlice  []string

	// Capture the capacity of the slice.
	sliceCap := cap(nilSlice)

	// Append ~100k strings to the slice.
	for record := 1; record <= 1e5; record++ {


		// Use the built-in function append to add to the slice.
		value := fmt.Sprintf("Rec: %d", record)
		 nilSlice  = append(nilSlice, value)

		// When the capacity of the slice changes, display the changes.
		if	sliceCap != cap(nilSlice){

			capChg := float64(cap(nilSlice) - sliceCap) / float64(sliceCap) * 100 

			sliceCap = cap(nilSlice)
			fmt.Printf("Address[ %p ]\t Index[ %d ] \t\t Cap[ %d - %2.f%% ] \n",
			&nilSlice[0],
			record,
			cap(nilSlice),
			capChg,
		)


		}

	}


}
