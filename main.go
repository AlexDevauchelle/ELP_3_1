package main

import (
	"container/heap"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

const dimension = 21
const waitTime = 1000

var ma_map [dimension][dimension]int

type Event struct {
	temps_event int
	genre       string
	origine     [2]int
	position    [2]int
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
	var event1 = Event{int(100), "depart", [2]int{1, 2}, [2]int{1, 2}, [2]int{3, 4}}
	var event2 = Event{int(20), "tirage", [2]int{3, 4}, [2]int{1, 2}, [2]int{1, 2}}
	var event3 = Event{int(50), "depart", [2]int{3, 4}, [2]int{1, 2}, [2]int{1, 2}}

	events := [3]Event{event1, event2, event3}
	// Create a priority queue, put the events in it, and
	pq := make(PriorityQueue, len(events))
	for i := 0; i < pq.Len(); i++ {
		pq[i] = &events[i]
	}
	heap.Init(&pq)
	// Insert a new event and then modify its priority.
	var event4 = Event{int(60), "depart", [2]int{3, 4}, [2]int{1, 2}, [2]int{1, 2}}
	heap.Push(&pq, &event4)
	pq.gestionTirage(int(20))
	// Take the items out; they arrive in increasing priority order.
	for pq.Len() > 0 {
		event := heap.Pop(&pq).(*Event)
		fmt.Println("--------------------------------------------------")
		fmt.Printf("l'évènement a lieu au tempss %d et est de genre %s agents", event.temps_event, event.genre)
		fmt.Println(".")
	}
}

func (pq *PriorityQueue) gestionTirage(temps_event int) {
	fmt.Printf("Temps %d : Un tirage à lieu !\n", temps_event)
	nouveauTirage := Event{temps_event + 30, "tirage", [2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}}
	heap.Push(pq, &nouveauTirage)

	proba := rand.Float64()
	lim_proba := 0.5 // est la proba d'avoir une intervention qui pop, à determiner en fonction de la fréquence des inter voulues

	// si la proba est inf à P(event), on instance une structure de classe event
	if proba < lim_proba {

		origine := [2]int{(dimension - 1) / 2, (dimension - 1) / 2}
		positionx := rand.Intn((dimension - 1))
		positiony := rand.Intn((dimension - 1))
		destination := [2]int{positionx, positiony}
		event := Event{temps_event + 2, "depart", origine, origine, destination}

		fmt.Printf("Temps %d : Un départ aura lieu au temps %d en destination de %d!\n", temps_event, temps_event+2, destination)
		// est ce qu'il serait judicieux de retourner un true quand il y a une intervention?

		heap.Push(pq, &event) // done : add event to heap

	}

}

func (pq *PriorityQueue) gestionDepart(depart Event) {

	deplacement := Event{temps_event: depart.temps_event + rand.Intn(5), genre: "deplacement", origine: depart.origine, position: depart.origine, destination: depart.destination}

	heap.Push(pq, &deplacement)
	fmt.Printf("Temps %d : Un agent part de %d sur une intervention en %d !\n", depart.temps_event, depart.origine, depart.destination)

}

func (pq *PriorityQueue) gestionDeplacement(pDeplacement Event) {

	x_or_y := rand.Intn(5)
	next_pos := [2]int{pDeplacement.position[0], pDeplacement.position[1]}
	if (x_or_y == 0) || (pDeplacement.position[1] == pDeplacement.destination[1]) {
		//Deplacement en X
		if pDeplacement.position[0] > pDeplacement.destination[0] {
			next_pos[0] -= 1
		} else {
			next_pos[0] += 1
		}

	} else {
		//Deplacement en Y
		if pDeplacement.position[1] > pDeplacement.destination[1] {
			next_pos[1] -= 1
		} else {
			next_pos[1] += 1
		}

	}

	if (next_pos[0] == pDeplacement.destination[0]) && (next_pos[1] == pDeplacement.destination[1]) { // Si la prochaine position est la destination alors soit arrive event soit retour event

		if ((dimension-1)/2 == pDeplacement.destination[0]) && ((dimension-1)/2 == pDeplacement.destination[1]) { //Si la destination est la centrale alors -> Arrive Event
			if (pDeplacement.position[0] != (dimension-1)/2) || (pDeplacement.position[1] != (dimension-1)/2) {
				ma_map[pDeplacement.position[0]][pDeplacement.position[1]] = 0 // On clear la position actuel de l'agent sur la map
			}

			arrive := Event{temps_event: pDeplacement.temps_event + rand.Intn(5), genre: "arrive", origine: pDeplacement.destination, position: pDeplacement.destination, destination: pDeplacement.origine}
			heap.Push(pq, &arrive)

		} else { // Sinon -> Retour Event
			if (pDeplacement.position[0] != (dimension-1)/2) || (pDeplacement.position[1] != (dimension-1)/2) { //Si la position actuel n'est pas la centrale
				ma_map[pDeplacement.position[0]][pDeplacement.position[1]] = 0 // On clear la position actuel de l'agent sur la map
			}
			if (next_pos[0] != (dimension-1)/2) && (next_pos[1] != (dimension-1)/2) { //Si la prochaine position n'est pas la centrale
				ma_map[pDeplacement.position[0]][pDeplacement.position[1]] = 0 // On clear la position actuel de l'agent sur la map
			}
			ma_map[next_pos[0]][next_pos[1]] = 3 //On occupe la nouvelle position de l'agent avec le code d'intervention => 3

			ma_map[pDeplacement.position[0]][pDeplacement.position[1]] = 0 // On clear la position actuel de l'agent sur la map

			retour := Event{temps_event: pDeplacement.temps_event + rand.Intn(5), genre: "retour", origine: pDeplacement.destination, position: pDeplacement.destination, destination: pDeplacement.origine}
			heap.Push(pq, &retour)
		}

	} else { //Sinon deplacement
		if (pDeplacement.position[0] != (dimension-1)/2) || (pDeplacement.position[1] != (dimension-1)/2) {
			ma_map[pDeplacement.position[0]][pDeplacement.position[1]] = 0 // On clear la position actuel de l'agent sur la map
		}
		ma_map[next_pos[0]][next_pos[1]] = 1 //On occupe la nouvelle position de l'agent avec le code de deplacement => 0

		nDeplacement := Event{temps_event: pDeplacement.temps_event + rand.Intn(5), genre: "deplacement", origine: pDeplacement.origine, position: next_pos, destination: pDeplacement.destination}
		heap.Push(pq, &nDeplacement)

	}
	fmt.Printf("Temps %d : Un agent en %d se deplace en %d !\n", pDeplacement.temps_event, pDeplacement.position, next_pos)

}

func (pq *PriorityQueue) gestionRetour(e Event) {

	fmt.Printf("Temps %d : Un agent est arrivé en %d sur l'intervention ! Il rentre en %d.\n", e.temps_event, e.origine, e.destination)

	deplacement := Event{temps_event: e.temps_event + rand.Intn(20), genre: "deplacement", origine: e.origine, position: e.origine, destination: e.destination}

	heap.Push(pq, &deplacement)

}

func (pq *PriorityQueue) gestionHeap() {

	for {
		time.Sleep(waitTime * time.Millisecond)

		if pq.Len() > 0 {

			event := heap.Pop(pq).(*Event)

			if event.genre == "tirage" {
				pq.gestionTirage(event.temps_event)
			}

			if event.genre == "depart" {
				pq.gestionDepart(*event)
			}

			if event.genre == "deplacement" {
				pq.gestionDeplacement(*event)
			}

			if event.genre == "retour" {
				pq.gestionRetour(*event)
			}
			if event.genre == "arrive" {
				fmt.Printf("Temps %d : Un agent est rentré à la centrale !\n", event.temps_event)
			}
			printMap()

		} else {
			return
		}

	}

}

func updateImage() {

	width := len(ma_map[0])
	height := len(ma_map)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	red := color.RGBA{255, 0, 0, 0xff}

	blue := color.RGBA{0, 0, 255, 0xff}

	// Set color for each pixel.
	for y := 0; y < len(ma_map); y++ {
		for x := 0; x < len(ma_map[0]); x++ {
			switch {
			case ma_map[y][x] == 0: // case vide
				img.Set(x, y, color.Black)
			case ma_map[y][x] == 1: // case deplacement
				img.Set(x, y, color.White)
			case ma_map[y][x] == 3: // case intervention
				img.Set(x, y, red)
			case ma_map[y][x] == 9: // case centrale
				img.Set(x, y, blue)
			default:
				// Use zero value.
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)

}

func printMap() {

	output := ""
	for y := 0; y < len(ma_map); y++ {
		for x := 0; x < len(ma_map[0]); x++ {
			switch {
			case ma_map[y][x] == 0: // case vide
				output += " "
			case ma_map[y][x] == 1: // case deplacement
				output += "1"
			case ma_map[y][x] == 3: // case intervention
				output += "X"
			case ma_map[y][x] == 9: // case centrale
				output += "C"
			default:
				// Use zero value.
			}

		}
		output += "\n"
	}
	fmt.Printf(output)
}

func managecollision() {
	//Au lieu de calculer un Dijkstra à chaque deplacement de case de l'agent, ce qui serait couteux en temps, on va implementer une fonction qui gère les collisions
	tentative_deplacement := 0

	//si la prochaine case indiquée par la fonction gestiondéplacement est dispo
	//if next_pos == 0 {
		//deplacement sur la prochaine case N+1
		//case N passe à 0
	}
	//si l'agent ne peut pas se deplacer
	if tentative_deplacement == 3 {
		fmt.Println("le plateau est chargé, l'agent fait une pause sur sa case.")
		time.Sleep(8 * time.Second) // on attend un nombre n de secondes pour que le plateau se libère

	} else {
		//run la fonction gestiondeplacement
		tentative_deplacement++
	}
}

func main() {
	ma_map[(dimension-1)/2][(dimension-1)/2] = 9

	//e := Event{temps_event: 10, genre: "tirage", origine: [2]int{0, 0}, position: [2]int{0, 0}, destination: [2]int{0, 0}}

	e := Event{temps_event: 10, genre: "depart", origine: [2]int{(dimension - 1) / 2, (dimension - 1) / 2}, position: [2]int{(dimension - 1) / 2, (dimension - 1) / 2}, destination: [2]int{((dimension - 1) / 2) + 5, ((dimension - 1) / 2) + 5}}

	tab_event := [1]Event{e}

	pq := make(PriorityQueue, len(tab_event))
	for i := 0; i < pq.Len(); i++ {
		pq[i] = &tab_event[i]
	}
	heap.Init(&pq)

	pq.gestionHeap()

}
