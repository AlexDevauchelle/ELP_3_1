package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

type Event struct {
	temps_event int
	genre       string
	origine     [2]int
	destination [2]int
}

type PriorityQueue []*Event //crée une liste d'éléments Event

func (pq PriorityQueue) Len() int { return len(pq) } //donne la taille de la file

func (pq PriorityQueue) Less(i, j int) bool { //compare les prioritées
	return pq[i].temps_event < pq[j].temps_event
}
func (pq PriorityQueue) Swap(i, j int) { //échange les éléments de rang i et j
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) { //ajoute un élément à sa place dans la liste
	event := x.(*Event)
	*pq = append(*pq, event)
}
func (pq *PriorityQueue) Pop() interface{} { //donne l'élément prioritaire de la liste et le fait sortir de celle ci
	old := *pq
	n := len(old)
	event := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return event
}
func (pq *PriorityQueue) update(event *Event, temp_event int) { //change le temps de trajet (possibilité de changer d'autre trucs)
	event.temps_event = temp_event
	heap.Fix(pq, event.temps_event)
}
func Example_priorityQueue() {
	var event1 = Event{int(100), "depart", [2]int{1, 2}, [2]int{3, 4}}
	var event2 = Event{int(20), "tirage", [2]int{3, 4}, [2]int{1, 2}}
	var event3 = Event{int(50), "depart", [2]int{3, 4}, [2]int{1, 2}}

	events := [3]Event{event1, event2, event3}
	// Create a priority queue, put the events in it, and
	pq := make(PriorityQueue, len(events))
	for i := 0; i < pq.Len(); i++ {
		pq[i] = &events[i]
	}
	heap.Init(&pq)
	// Insert a new event and then modify its priority.
	var event4 = Event{int(60), "depart", [2]int{3, 4}, [2]int{1, 2}}
	heap.Push(&pq, &event4)
	pq.newEvent(int(20))
	// Take the items out; they arrive in increasing priority order.
	for pq.Len() > 0 {
		event := heap.Pop(&pq).(*Event)
		fmt.Println("--------------------------------------------------")
		fmt.Printf("l'évènement a lieu au tempss %d et est de genre %s agents", event.temps_event, event.genre)
		fmt.Println(".")
	}
}

func (pq *PriorityQueue) newEvent(temps_event int) {
	fmt.Printf("Temps %d : Un tirage à lieu !\n", temps_event)
	nouveauTirage := Event{temps_event + 5, "tirage", [2]int{0, 0}, [2]int{0, 0}}
	heap.Push(pq, &nouveauTirage)

	proba := rand.Float64()
	lim_proba := 0.5 // est la proba d'avoir une intervention qui pop, à determiner en fonction de la fréquence des inter voulues

	// si la proba est inf à P(event), on instance une structure de classe event
	if proba < lim_proba {

		origine := [2]int{0, 0}
		positionx := rand.Intn(500)
		positiony := rand.Intn(500)
		destination := [2]int{positionx, positiony}
		event := Event{temps_event + 2, "depart", origine, destination}

		fmt.Printf("Temps %d : Un départ aura lieu au temps %d en destination de %d!\n", temps_event, temps_event+2, destination)
		// est ce qu'il serait judicieux de retourner un true quand il y a une intervention?

		heap.Push(pq, &event) // done : add event to heap

	}

}

func (pq *PriorityQueue) newRetour(depart Event) {

	retour := Event{temps_event: depart.temps_event + rand.Intn(20), genre: "retour", origine: depart.origine, destination: depart.destination}

	heap.Push(pq, &retour)
	fmt.Printf("Temps %d : Un agent part de %d sur une intervention en %d !\n", depart.temps_event, depart.origine, depart.destination)

}

func (pq *PriorityQueue) newArrive(retour Event) {

	fmt.Printf("Temps %d : Un agent est arrivé en %d sur l'intervention ! Il rentre en %d.\n", retour.temps_event, retour.destination, retour.origine)

	arrive := Event{temps_event: retour.temps_event + rand.Intn(20), genre: "arrive", origine: retour.destination, destination: retour.origine}

	heap.Push(pq, &arrive)

}

func (pq *PriorityQueue) gestionHeap() {

	for {
		time.Sleep(1000 * time.Millisecond)

		event := heap.Pop(pq).(*Event)

		if event.genre == "tirage" {
			pq.newEvent(event.temps_event)
		}

		if event.genre == "depart" {
			pq.newRetour(*event)
		}

		if event.genre == "retour" {
			pq.newArrive(*event)
		}
		if event.genre == "arrive" {
			fmt.Printf("Temps %d : Un agent est rentré à la centrale !\n", event.temps_event)
		}

	}

}

func main() {
	e := Event{temps_event: 10, genre: "tirage", origine: [2]int{0, 0}, destination: [2]int{0, 0}}
	tab_event := [1]Event{e}

	pq := make(PriorityQueue, len(tab_event))
	for i := 0; i < pq.Len(); i++ {
		pq[i] = &tab_event[i]
	}
	heap.Init(&pq)

	pq.gestionHeap()

}
