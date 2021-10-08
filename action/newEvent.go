package main

import (
	"fmt"
	"math/rand"
)

type Event struct {
	proba			float32
	lim_proba		float32
	temp_event		float32
	origine         [2]int
	localisation    [2]int
	agents_required	[2]int	// le nombre d'agents demandés pour l'intervention
}
func newEvent(proba float32, temp_event float32, origine [2]int, localisation [2]int, agents_required [2]int ) {
	proba = rand.Intn(1)
	lim_proba = 0.5 // est la proba d'avoir une intervention qui pop, à determiner en fonction de la fréquence des inter voulues
	
	// si la proba est inf à P(event), on instance une structure de classe event
	if proba < lim_proba {
		Event {
			event := Event{origine, localisation, temps_event, agents_required}
			fmt.Println("c'est bon")
		return event// est ce qu'il serait judicieux de retourner un true quand il y a une intervention?
		}
	}
	}