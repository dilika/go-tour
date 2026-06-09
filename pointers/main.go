package main 


import "fmt"


func main(){

	count := 10

	// Display variable count address and value 
	
	fmt.Println("Value of [",count, "] and address of [", &count ,"] ")

	incremente1(count)
	fmt.Println("Value of [",count, "] and address of [", &count ,"] ")
	incremente2(&count)




}

func incremente1( newCount int ) {
		newCount++

	fmt.Println("Value of incremente1  newCount [",newCount, "] and address of [", &newCount ,"] ")
	}
	
func incremente2( newCount *int ) {
		*newCount++

	fmt.Println("Value of  incremente2 newCount [",newCount, "] and address of [", &newCount ,"] ")
	}

