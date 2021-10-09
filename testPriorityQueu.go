package main

import (
	"action"
	"container/heap"
)
type PriorityQueue []*arrive.Arrive //crée une liste d'éléments arrive
func (pq PriorityQueue) Len() int { return len(pq) }//donne la taille de la file

func (pq PriorityQueue) Less(i, j int) bool {//compare les prioritées
	return pq[i].Temps_trajet < pq[j].Temps_trajet
}

func (pq PriorityQueue) Swap(i, j int) {//échange les éléments de rang i et j 
	pq[i], pq[j] = pq[j], pq[i]

}

func (pq *PriorityQueue) Push(x interface{}) {//ajoute un élément à sa place dans la liste
	arrive1 := x.(*arrive.Arrive)
	*pq = append(*pq, arrive1)
}
func (pq *PriorityQueue) Pop() interface{} {//donne l'élément prioritaire de la liste et le fait sortir de celle ci
	old := *pq
	n := len(old)
	arrive1 := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*pq = old[0 : n-1]
	return arrive1
}
func (pq *PriorityQueue) update(arrive1 *arrive.Arrive, temps_trajet int) {//change le temps de trajet (possibilité de changer d'autre trucs)
	arrive1.Temps_trajet = temps_trajet
	heap.Fix(pq, arrive1.Temps_trajet)

}

func Example_priorityQueue() {
	var mavar = arrive.New([2]int{1, 2}, [2]int{3, 4}, int(16), true)
	var mavar2 = arrive.New([2]int{3, 4}, [2]int{1, 2}, int(193), false)
	var mavar3 = arrive.New([2]int{3, 4}, [2]int{1, 2}, int(112), false)
	var mavar4 = arrive.New([2]int{3, 4}, [2]int{1, 2}, int(190), false)
	var mavar5 = arrive.New([2]int{3, 4}, [2]int{1, 2}, int(130), false)
	arrivees := [5]arrive.Arrive{mavar,mavar2,mavar3,mavar4,mavar5}

	// Create a priority queue, put the events in it, and
	pq := make(PriorityQueue, len(arrivees))
	
	for i := 0;i<pq.Len();i++  {
		pq[i] = &arrivees[i]	
	}
	heap.Init(&pq)

	// Insert a new event and then modify its priority.	
	var mavar6 = arrive.New([2]int{3, 4}, [2]int{1, 2}, int(160), false)

	heap.Push(&pq, &mavar6)

	// Take the items out; they arrive in increasing priority order.
	for pq.Len() > 0 {
		arrive1 := heap.Pop(&pq).(*arrive.Arrive)
		arrive1.Affichage()
	}
}

func main() {
	Example_priorityQueue()
}

