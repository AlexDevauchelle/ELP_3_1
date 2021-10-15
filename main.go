package main

import (
	"container/heap"
	"fmt"
	"math/rand"
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
	return pq[i].temp_event < pq[j].temp_event
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
	event.temp_event = temp_event
	heap.Fix(pq, event.temp_event)
}
func Example_priorityQueue() {
	var event1 = Event{int(100),"depart",[2]int{1, 2}, [2]int{3, 4}}
	var event2 = Event{ int(20),"tirage",[2]int{3, 4}, [2]int{1, 2}}
	var event3 = Event{int(50),"depart", [2]int{3, 4}, [2]int{1, 2}}

	events := [3]Event{event1,event2,event3}
	// Create a priority queue, put the events in it, and
	pq := make(PriorityQueue, len(events))
	for i := 0;i<pq.Len();i++  {
		pq[i] = &events[i]	
	}
	heap.Init(&pq)
	// Insert a new event and then modify its priority.	
	var event4 = Event{int(60),"depart",[2]int{3, 4}, [2]int{1, 2}}
	heap.Push(&pq, &event4)
	pq.newEvent(int(20))
	// Take the items out; they arrive in increasing priority order.
	for pq.Len() > 0 {
		event := heap.Pop(&pq).(*Event)
		fmt.Println("--------------------------------------------------")
		fmt.Printf("l'évènement a lieu au tempss %d et est de genre %d agents",event.temps_event,event.genre)
		fmt.Println(".")
	}
}

func (pq *PriorityQueue) newEvent(temps_event int) {
	nouveauTirage := Event{temps_event+5,"tirage",[2]int{0,0},[2]int{0,0}}
	heap.Push(pq, &nouveauTirage)
	proba := rand.Float64()
	lim_proba := 0.5 // est la proba d'avoir une intervention qui pop, à determiner en fonction de la fréquence des inter voulues

	// si la proba est inf à P(event), on instance une structure de classe event
	if proba < lim_proba {
		origine :=[2]int{0,0}
		positionx := rand.Intn(500)
		positiony := rand.Intn(500)
		localisation:=[2]int{positionx,positiony}
		event := Event{temps_event+2,"depart", origine, localisation}

		fmt.Println("c'est bon")
		// est ce qu'il serait judicieux de retourner un true quand il y a une intervention?

		fmt.Println(event.origine)

		heap.Push(pq, &event)// done : add event to heap

	}

}

func main() {
	fmt.Println("Hello world")

	e := Event{temps_event: 10, genre: "tirage", origine: [2]int{0, 0}, destination: [2]int{0, 0}}

	Example_priorityQueue()

}
