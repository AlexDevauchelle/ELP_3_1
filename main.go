package main

import (
	"fmt"
	"math/rand"
	"container/heap"
)

type Arrive struct {
	origine         [2]int
	localisation    [2]int
	temps_trajet    int
	is_Intervention bool
}

type Depart struct {
	origine         [2]int
	destination     [2]int
	temps_trajet    int
	is_Intervention bool
}

type Event struct {
	temp_event      int
	origine         [2]int
	localisation    [2]int
	agents_required int // le nombre d'agents demandés pour l'intervention
}

type PriorityQueue []*Event //crée une liste d'éléments Event
func (pq PriorityQueue) Len() int { return len(pq) }//donne la taille de la file

func (pq PriorityQueue) Less(i, j int) bool {//compare les prioritées
	return pq[i].temp_event < pq[j].temp_event
}
func (pq PriorityQueue) Swap(i, j int) {//échange les éléments de rang i et j 
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {//ajoute un élément à sa place dans la liste
	event := x.(*Event)
	*pq = append(*pq, event)
}
func (pq *PriorityQueue) Pop() interface{} {//donne l'élément prioritaire de la liste et le fait sortir de celle ci
	old := *pq
	n := len(old)
	event := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*pq = old[0 : n-1]
	return event
}
func (pq *PriorityQueue) update(event *Event, temp_event int) {//change le temps de trajet (possibilité de changer d'autre trucs)
	event.temp_event = temp_event
	heap.Fix(pq, event.temp_event)
}
func Example_priorityQueue() {
	var event1 = Event{int(100),[2]int{1, 2}, [2]int{3, 4}, int(2)}
	var event2 = Event{ int(20),[2]int{3, 4}, [2]int{1, 2}, int(3)}
	var event3 = Event{int(50), [2]int{3, 4}, [2]int{1, 2}, int(2)}

	events := [3]Event{event1,event2,event3}
	// Create a priority queue, put the events in it, and
	pq := make(PriorityQueue, len(events))
	for i := 0;i<pq.Len();i++  {
		pq[i] = &events[i]	
	}
	heap.Init(&pq)
	// Insert a new event and then modify its priority.	
	var event4 = Event{int(60),[2]int{3, 4}, [2]int{1, 2}, int(1)}
	heap.Push(&pq, &event4)
	pq.newEvent(int(160),[2]int{3, 4}, int(6))
	// Take the items out; they arrive in increasing priority order.
	for pq.Len() > 0 {
		event := heap.Pop(&pq).(*Event)
		fmt.Println("--------------------------------------------------")
		fmt.Printf("l'évènement a lieu au temps %d et requiert %d agents",event.temp_event,event.agents_required)
		fmt.Println(".")
	}
}

func (pq *PriorityQueue) newEvent(temp_event int, origine [2]int, agents_required int) {
	proba := rand.Float64()
	lim_proba := 0.5 // est la proba d'avoir une intervention qui pop, à determiner en fonction de la fréquence des inter voulues

	// si la proba est inf à P(event), on instance une structure de classe event
	if proba < lim_proba {
		positionx := rand.Intn(500)
		positiony := rand.Intn(500)
		localisation:=[2]int{positionx,positiony}
		event := Event{temp_event, origine, localisation, agents_required}

		fmt.Println("c'est bon")
		// est ce qu'il serait judicieux de retourner un true quand il y a une intervention?

		fmt.Println(event.origine)

		heap.Push(pq, &event)// done : add event to heap

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
	maArrive := Arrive{[2]int{1, 2}, [2]int{3, 4}, int(3), true}
	maDepart := Depart{[2]int{1, 2}, [2]int{3, 4}, int(4), true}

	affichageArrive(maArrive)
	affichageDepart(maDepart)
/*
	for {
		newEvent(5.0, [2]int{1, 2}, [2]int{1, 2}, 5)
	}
*/
	Example_priorityQueue()

}
