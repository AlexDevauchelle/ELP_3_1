package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"
)

func getArgs() int {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: go run server.go <portnumber>\n")
		os.Exit(1)
	} else {
		fmt.Printf("#DEBUG ARGS Port Number : %s\n", os.Args[1])
		portNumber, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Usage: go run server.go <portnumber>\n")
			os.Exit(1)
		} else {
			return portNumber
		}

	}
	return -1
}

func main() {
	port := getArgs()
	fmt.Printf("#DEBUG MAIN Creating TCP Server on port %d\n", port)
	portString := fmt.Sprintf(":%s", strconv.Itoa(port))
	fmt.Printf("#DEBUG MAIN PORT STRING |%s|\n", portString)

	ln, err := net.Listen("tcp", portString)
	if err != nil {
		fmt.Printf("#DEBUG MAIN Could not create listener\n")
		panic(err)
	}

	//If we're here, we did not panic and ln is a valid listener

	connum := 1

	for {
		fmt.Printf("#DEBUG MAIN Accepting next connection\n")
		conn, errconn := ln.Accept()

		if errconn != nil {
			fmt.Printf("DEBUG MAIN Error when accepting next connection\n")
			panic(errconn)

		}

		//If we're here, we did not panic and conn is a valid handler to the new connection

		go handleConnection(conn, connum)
		connum += 1

	}
}

func handleConnection(connection net.Conn, connum int) {

	defer connection.Close()
	connReader := bufio.NewReader(connection)
	//    if err !=nil{
	//        fmt.Printf("#DEBUG %d handleConnection could not create reader\n", connum)
	//        return
	//    }

	/*
		variables, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Printf("#DEBUG %d RCV ERROR no panic, just a client\n", connum)
			fmt.Printf("Error :|%s|\n", err.Error())
		}

		variables = strings.TrimSuffix(variables, "\n")
		fmt.Printf("%s\n", variables)
		time.Sleep(10000 * time.Millisecond)
		tab_variables := strings.Split(variables, ",")

		taille_map := tab_variables[0]
		temps_sim := tab_variables[1]

	*/

	taille_map, err1 := connReader.ReadString('\n')
	if err1 != nil {
		fmt.Printf("#DEBUG %d RCV ERROR no panic, just a client\n", connum)
		fmt.Printf("Error :|%s|\n", err1.Error())
	}

	temps_sim, err2 := connReader.ReadString('\n')
	if err2 != nil {
		fmt.Printf("#DEBUG %d RCV ERROR no panic, just a client\n", connum)
		fmt.Printf("Error :|%s|\n", err2.Error())
	}

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	tm := 0
	ts := 0

	submatchall := re.FindAllString(taille_map, -1)
	for _, element := range submatchall {

		tm, _ = strconv.Atoi(element)

	}

	submatchall2 := re.FindAllString(temps_sim, -1)
	for _, element := range submatchall2 {

		ts, _ = strconv.Atoi(element)

	}

	/*
		ts, errts := strconv.Atoi(temps_sim)
		if errts != nil {
			fmt.Printf("#DEBUG %d temps_sim ne peut pas être convertit en interger\n", connum)
			fmt.Printf("Error :|%s|\n", errts.Error())
		}*/
	clientConn = connection

	fmt.Printf("Début d'une simulation sur une map de taille %d avec un temps max de %d\n", tm, ts)
	io.WriteString(connection, fmt.Sprintf("Début d'une simulation sur une map de taille %d avec un temps max de %d\n$", tm, ts))
	start_simu(tm, ts)

	io.WriteString(connection, fmt.Sprintf("End_Of_Connection$"))

}

var clientConn net.Conn

var dimension = 21

var time_limit = 250

const waitTime = 50

var ma_map [][]int

type Event struct {
	temps_event           int
	genre                 string
	origine               [2]int
	position              [2]int
	destination           [2]int
	tentative_deplacement int
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

func (pq *PriorityQueue) gestionTirage(temps_event int) {
	io.WriteString(clientConn, fmt.Sprintf("Temps %d : Un tirage à lieu !\n$", temps_event))

	//fmt.Printf("Temps %d : Un tirage à lieu !\n", temps_event)
	nouveauTirage := Event{temps_event + 5, "tirage", [2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}, 0}
	heap.Push(pq, &nouveauTirage)

	proba := rand.Float64()
	lim_proba := 0.7 // est la proba d'avoir une intervention qui pop, à determiner en fonction de la fréquence des inter voulues

	// si la proba est inf à P(event), on instance une structure de classe event
	if proba < lim_proba {

		origine := [2]int{(dimension - 1) / 2, (dimension - 1) / 2}
		positionx := rand.Intn((dimension - 1))
		positiony := rand.Intn((dimension - 1))
		if (dimension-1)/2 != positionx || ((dimension-1)/2 != positiony) {
			destination := [2]int{positionx, positiony}
			event := Event{temps_event + 2, "depart", origine, origine, destination, 0}
			io.WriteString(clientConn, fmt.Sprintf("Temps %d : Un départ aura lieu au temps %d en destination de %d!\n$", temps_event, temps_event+2, destination))

			//fmt.Printf("Temps %d : Un départ aura lieu au temps %d en destination de %d!\n", temps_event, temps_event+2, destination)
			heap.Push(pq, &event) //  add event to heap
		}

	}

}

func (pq *PriorityQueue) gestionDepart(depart Event) {

	deplacement := Event{temps_event: depart.temps_event + rand.Intn(5), genre: "deplacement", origine: depart.origine, position: depart.origine, destination: depart.destination, tentative_deplacement: 0}

	heap.Push(pq, &deplacement)

	io.WriteString(clientConn, fmt.Sprintf("Temps %d : Un agent part de %d sur une intervention en %d !\n$", depart.temps_event, depart.origine, depart.destination))
	//fmt.Printf("Temps %d : Un agent part de %d sur une intervention en %d !\n", depart.temps_event, depart.origine, depart.destination)

}

func (pq *PriorityQueue) gestionDeplacement(pDeplacement Event, next_pos [2]int) {
	if ((dimension-1)/2 != pDeplacement.position[0]) || ((dimension-1)/2 != pDeplacement.position[1]) { //si l'agent n'est pas à la centrale à l'origine
		ma_map[pDeplacement.position[0]][pDeplacement.position[1]] = 0
	}
	if (next_pos[0] == pDeplacement.destination[0]) && (next_pos[1] == pDeplacement.destination[1]) { // Si la prochaine position est la destination alors soit arrive event soit retour event
		if ((dimension-1)/2 == pDeplacement.destination[0]) && ((dimension-1)/2 == pDeplacement.destination[1]) { //Si la destination est la centrale alors -> Arrive Event
			arrive := Event{temps_event: pDeplacement.temps_event + rand.Intn(5), genre: "arrive", origine: pDeplacement.destination, position: pDeplacement.destination, destination: pDeplacement.origine, tentative_deplacement: 0}
			heap.Push(pq, &arrive)
		} else { // Sinon -> Retour Event
			ma_map[next_pos[0]][next_pos[1]] = 3 //On occupe la nouvelle position de l'agent avec le code d'intervention => 3
			retour := Event{temps_event: pDeplacement.temps_event + rand.Intn(5), genre: "retour", origine: pDeplacement.destination, position: pDeplacement.destination, destination: pDeplacement.origine, tentative_deplacement: 0}
			heap.Push(pq, &retour)
		}

	} else { //Sinon deplacement
		ma_map[next_pos[0]][next_pos[1]] = 1 //On occupe la nouvelle position de l'agent avec le code de deplacement => 0
		nDeplacement := Event{temps_event: pDeplacement.temps_event + rand.Intn(5), genre: "deplacement", origine: pDeplacement.origine, position: next_pos, destination: pDeplacement.destination, tentative_deplacement: 0}
		heap.Push(pq, &nDeplacement)
	}
	io.WriteString(clientConn, fmt.Sprintf("Temps %d : Un agent en %d se deplace en %d !\n$", pDeplacement.temps_event, pDeplacement.position, next_pos))
	//fmt.Printf("Temps %d : Un agent en %d se deplace en %d !\n", pDeplacement.temps_event, pDeplacement.position, next_pos)
}

func (pq *PriorityQueue) managecollision(pDeplacement Event) {
	//Au lieu de calculer un Dijkstra à chaque deplacement de case de l'agent, ce qui serait couteux en temps, on va implementer une fonction qui gère les collisions
	destination := [2]int{pDeplacement.destination[0], pDeplacement.destination[1]}
	deplacement_vers_objectif := true
	xActuel := pDeplacement.position[0]
	yActuel := pDeplacement.position[1]

	//regarde les alentours
	north := false
	south := false
	east := false
	west := false
	if yActuel < dimension-1 {
		south = ma_map[xActuel][yActuel+1] == 0 || ma_map[xActuel][yActuel+1] == 9
	}
	if yActuel > 0 {
		north = ma_map[xActuel][yActuel-1] == 0 || ma_map[xActuel][yActuel-1] == 9
	}
	if xActuel < dimension-1 {
		east = ma_map[xActuel+1][yActuel] == 0 || ma_map[xActuel+1][yActuel] == 9
	}
	if xActuel > 0 {
		west = ma_map[xActuel-1][yActuel] == 0 || ma_map[xActuel-1][yActuel] == 9
	}
	next_pos := [2]int{xActuel, yActuel}
	x_or_y := rand.Intn(2)
	if pDeplacement.tentative_deplacement > 3 { //on regarde si on est pas bloqué depuis trop longtemps
		if yActuel < destination[1] && north {
			destination[1] = yActuel - 1
		} else if south {
			destination[1] = yActuel + 1
		}
		if xActuel < destination[0] && east {
			destination[0] = xActuel - 1
		} else if west {
			destination[0] = xActuel + 1
		}
		deplacement_vers_objectif = false
		if rand.Intn(100) < 33 { //rends plus ou moins aléatoire le nombre de fois où l'agent va reculer
			pDeplacement.tentative_deplacement = 0
		}
	}
	//essaye de se déplacer
	if x_or_y == 1 {
		if yActuel == destination[1] {
			//Deplacement en X
			if xActuel > destination[0] && west {
				next_pos[0] -= 1
			} else if xActuel < destination[0] && east {
				next_pos[0] += 1
			}
		} else {
			//Deplacement en Y
			if yActuel > destination[1] && north {
				next_pos[1] -= 1
			} else if yActuel < destination[1] && south {
				next_pos[1] += 1
			}
		}
	} else {
		if xActuel == destination[0] {
			//Deplacement en Y
			if yActuel > destination[1] && south {
				next_pos[1] -= 1
			} else if yActuel < destination[1] && north {
				next_pos[1] += 1
			}

		} else {
			//Deplacement en X
			if xActuel > destination[0] && west {
				next_pos[0] -= 1
			} else if xActuel < destination[0] && east {
				next_pos[0] += 1
			}
		}
	}
	if xActuel == next_pos[0] && yActuel == next_pos[1] {
		deplacement := Event{pDeplacement.temps_event + rand.Intn(5), "deplacement", pDeplacement.origine, pDeplacement.position, pDeplacement.destination, pDeplacement.tentative_deplacement + 1}
		heap.Push(pq, &deplacement)
	} else {
		if deplacement_vers_objectif {
			pDeplacement.tentative_deplacement = 0
		}
		pq.gestionDeplacement(pDeplacement, next_pos)
	}

}

func (pq *PriorityQueue) gestionRetour(e Event) {
	io.WriteString(clientConn, fmt.Sprintf("Temps %d : Un agent est arrivé en %d sur l'intervention ! Il rentre en %d.\n$", e.temps_event, e.origine, e.destination))
	//fmt.Printf("Temps %d : Un agent est arrivé en %d sur l'intervention ! Il rentre en %d.\n", e.temps_event, e.origine, e.destination)

	deplacement := Event{temps_event: e.temps_event + rand.Intn(20), genre: "deplacement", origine: e.origine, position: e.origine, destination: e.destination, tentative_deplacement: 0}

	heap.Push(pq, &deplacement)

}

func (pq *PriorityQueue) gestionHeap() {

	for {
		time.Sleep(waitTime * time.Millisecond)

		if pq.Len() > 0 {

			event := heap.Pop(pq).(*Event)

			if event.temps_event >= time_limit {
				io.WriteString(clientConn, fmt.Sprintf("Temps %d : Simulation terminée !\n$", time_limit))
				fmt.Printf("Temps %d : Simulation terminée !\n", time_limit)
				printMap()
				break
			}

			if event.genre == "tirage" {
				pq.gestionTirage(event.temps_event)
			}

			if event.genre == "depart" {
				pq.gestionDepart(*event)
			}

			if event.genre == "deplacement" {
				pq.managecollision(*event)
			}

			if event.genre == "retour" {
				pq.gestionRetour(*event)
			}
			if event.genre == "arrive" {
				io.WriteString(clientConn, fmt.Sprintf("Temps %d : Un agent est rentré à la centrale !\n$", event.temps_event))
				//fmt.Printf("Temps %d : Un agent est rentré à la centrale !\n", event.temps_event)
			}
			printMap()

		} else {
			return
		}

	}

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
	io.WriteString(clientConn, fmt.Sprintf("%s$", output))
	//fmt.Printf(output)
}

func start_simu(taille_map int, temps_simu int) {
	dimension = taille_map
	time_limit = temps_simu

	for y := 0; y < dimension; y++ {
		ma_map = append(ma_map, []int{})
		for x := 0; x < dimension; x++ {
			ma_map[y] = append(ma_map[y], 0)

		}
	}

	ma_map[(dimension-1)/2][(dimension-1)/2] = 9

	e := Event{temps_event: 1, genre: "tirage", origine: [2]int{0, 0}, position: [2]int{0, 0}, destination: [2]int{0, 0}, tentative_deplacement: 0}

	//e := Event{temps_event: 10, genre: "depart", origine: [2]int{(dimension - 1) / 2, (dimension - 1) / 2}, position: [2]int{(dimension - 1) / 2, (dimension - 1) / 2}, destination: [2]int{((dimension - 1) / 2) + 5, ((dimension - 1) / 2) + 5}, tentative_deplacement : 0}

	tab_event := [1]Event{e}

	pq := make(PriorityQueue, len(tab_event))
	for i := 0; i < pq.Len(); i++ {
		pq[i] = &tab_event[i]
	}
	heap.Init(&pq)
	pq.gestionHeap()

}
