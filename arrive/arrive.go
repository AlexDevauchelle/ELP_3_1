package arrive

import (
	"fmt"
)

type Arrive struct {
	origine         [2]int
	localisation    [2]int
	temps_trajet    float32
	is_Intervention bool
}

/*
Créer une instance de la classe Arrive

@return: struct Arrive
*/
func New(origine [2]int, localisation [2]int, temps_trajet float32, is_Intervention bool) Arrive {
	arrive := Arrive{origine, localisation, temps_trajet, is_Intervention}
	return arrive
}

/*
Affiche des informations sur un Arrive

@return: void
*/
func (a Arrive) Affichage() { // déclaration de ma méthode Affichage() liée à ma structure Arrive
	fmt.Println("--------------------------------------------------")
	fmt.Println("Origine", a.origine)
	fmt.Println("Localisation", a.localisation)

	if a.is_Intervention {
		fmt.Println("C'est une intervention")
	} else {
		fmt.Println("Ce n'est pas une intervention")
	}

	fmt.Println("\nLe temps de trajet est de ", a.temps_trajet)

}
