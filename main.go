package main

import (
	"fmt"
	"math/rand"
)

type Arrive struct {
	origine         [2]int
	localisation    [2]int
	temps_trajet    float32
	is_Intervention bool
}

type Depart struct {
	origine         [2]int
	destination     [2]int
	temps_trajet    float32
	is_Intervention bool
}

type Event struct {
	temp_event      float32
	origine         [2]int
	localisation    [2]int
	agents_required int // le nombre d'agents demandés pour l'intervention
}

func newEvent(temp_event float32, origine [2]int, localisation [2]int, agents_required int) {
	proba := rand.Float64()
	lim_proba := 0.5 // est la proba d'avoir une intervention qui pop, à determiner en fonction de la fréquence des inter voulues

	// si la proba est inf à P(event), on instance une structure de classe event
	if proba < lim_proba {
		event := Event{temp_event, origine, localisation, agents_required}
		fmt.Println("c'est bon")
		// est ce qu'il serait judicieux de retourner un true quand il y a une intervention?

		fmt.Println(event.origine)

		// TODO : add event to heap

	}

}

func affichageArrive(a Arrive) { // déclaration de ma méthode Affichage() liée à ma structure Arrive
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

func affichageDepart(d Depart) { // déclaration de ma méthode Affichage() liée à ma structure Depart
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

func main() {
	fmt.Println("Hello world")
	maArrive := Arrive{[2]int{1, 2}, [2]int{3, 4}, float32(3.3), true}
	maDepart := Depart{[2]int{1, 2}, [2]int{3, 4}, float32(3.3), true}

	affichageArrive(maArrive)
	affichageDepart(maDepart)

	for {
		newEvent(5.0, [2]int{1, 2}, [2]int{1, 2}, 5)
	}

}
