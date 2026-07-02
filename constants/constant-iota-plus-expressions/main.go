package main 


import "fmt"


func main() {

	const (
			_ = iota 
			KB = 1 << (10 * iota)
			MB
			GB
			TB	
	)

	fmt.Println(KB, MB, GB, TB)
}
