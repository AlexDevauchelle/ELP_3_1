package action

import (
	"fmt"
	"math/rand"
)

type Depart struct {
	origine         [2]int
	destination     [2]int
	temps_trajet    float32
	is_Intervention bool
}

/*
Créer une instance de la classe Depart

@return: struct Depart
*/
func New(origine [2]int, destination [2]int, is_Intervention bool) Depart {
	depart := Depart{origine, destination, rand.Float32(), is_Intervention}
	return depart
}

/*
Affiche des informations sur un Depart

@return: void
*/
func (d Depart) Affichage() { // déclaration de ma méthode Affichage() liée à ma structure Depart
	fmt.Println("--------------------------------------------------")
	fmt.Println("Origine", d.origine)
	fmt.Println("destination", d.destination)

	if d.is_Intervention {
		fmt.Println("C'est une intervention")
	} else {
		fmt.Println("Ce n'est pas une intervention")
	}

	fmt.Println("\nLe temps de trajet est de ", d.temps_trajet)

}
