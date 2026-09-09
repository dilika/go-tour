package main

import "fmt"

// Démo : le cours Ardan (Go tour, arrays) affirme que l'adresse de la variable
// de boucle `v` est IDENTIQUE à chaque itération. Vérifions sur Go 1.26.

func main() {
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	fmt.Println("Cours Ardan : \"&v a la meme adresse a chaque iteration\"")
	fmt.Println("Realite Go 1.26 :")
	fmt.Println()

	for i, v := range friends {
		fmt.Printf("iter %d  &i=%p  &v=%p  &friends[i]=%p\n", i, &i, &v, &friends[i])
	}

	fmt.Println()
	fmt.Println("&i  et &v  changent a chaque tour  -> variable fraiche par iteration")
	fmt.Println("&friends[i] avance de 16 octets    -> tableau contigu (string = 2 mots = 16B)")
	fmt.Println()
	fmt.Println("Cause : Go 1.22 a rendu la variable de boucle nouvelle a chaque iteration.")
	fmt.Println("C'est le MEME changement qui a corrige le piege closure en boucle (3 3 3 -> 0 1 2).")
}
