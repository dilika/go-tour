package main

import "fmt"


const (
	Read = 1 << iota
	Write
	Exec
)


func main(){

	fmt.Println("============== Flags Values =======================")
	fmt.Printf("Read = %d (binary: %04b)\n", Read, Read)
	fmt.Printf("Write = %d (binary: %04b)\n", Write, Write)
	fmt.Printf("Exec = %d (binary: %04b)\n", Exec, Exec)

	fmt.Println("\n============= Combine permission (with |) ========")
	permission := Read | Write
	fmt.Printf("Read + Write = %d (binary: %04b)\n", permission, permission)
  fullAccess := Read | Write | Exec
	fmt.Printf("fullAccess = %d (binary: %04b)\n", fullAccess, fullAccess)
	fmt.Println("\n============== Verify permission =================")
	fmt.Printf("Do I have Read permission? %v\n", permission&Read != 0)
	fmt.Printf("Do I have Write permission? %v\n", permission&Write != 0)
	fmt.Printf("Do I have Exec permission? %v\n", permission&Exec != 0)

	fmt.Println("=\n=============== Display Fonction ===================") printPermsission(permission)
	printPermsission(fullAccess)

}

func printPermsission(perm int) {
	fmt.Printf("Permission (%d) : ", perm)
	if perm&Read != 0 {
		fmt.Print("Read ")
	}

	if perm&Write != 0 {
		fmt.Print("Write ")
	}

	if perm&Exec != 0 {
		fmt.Print("Exec ")
	}

	fmt.Println()
}

