package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func getArgs() int {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: go run client.go <portnumber>\n")
		os.Exit(1)
	} else {
		fmt.Printf("#DEBUG ARGS Port Number : %s\n", os.Args[1])
		portNumber, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Usage: go run client.go <portnumber>\n")
			os.Exit(1)
		} else {
			return portNumber
		}

	}
	//Should never be reached
	return -1
}

func main() {
	port := getArgs()
	fmt.Printf("#DEBUG DIALING TCP Server on port %d\n", port)
	portString := fmt.Sprintf("127.0.0.1:%s", strconv.Itoa(port))
	fmt.Printf("#DEBUG MAIN PORT STRING |%s|\n", portString)

	conn, err := net.Dial("tcp", portString)
	if err != nil {
		fmt.Printf("#DEBUG MAIN could not connect\n")
		os.Exit(1)
	} else {

		defer conn.Close()
		reader := bufio.NewReader(conn)
		//fmt.Printf("#DEBUG MAIN connected\n")

		input := bufio.NewReader(os.Stdin)

		fmt.Printf("Quelle est la taille de la carte ?\n") // Question to the user
		taille_map, err1 := input.ReadString('\n')
		if err1 != nil {
			fmt.Printf("DEBUG MAIN could not read taille_map from stdIn")
			os.Exit(1)
		}
		//taille_map = strings.TrimSuffix(taille_map, "\n")

		io.WriteString(conn, fmt.Sprintf(taille_map))

		fmt.Printf("Quelle est la durée de la simulation ?\n") // Question to the user
		temps_sim, err2 := input.ReadString('\n')
		if err2 != nil {
			fmt.Printf("DEBUG MAIN could not read temps_sim from stdIn")
			os.Exit(1)
		}
		//temps_sim = strings.TrimSuffix(temps_sim, "\n")

		//tabstring := []string{taille_map, temps_sim}

		//varString := strings.Join(tabstring, ",")
		//fmt.Printf(varString + "\n")

		io.WriteString(conn, fmt.Sprintf(temps_sim))

		resultString, err := reader.ReadString('$')
		if err != nil {
			fmt.Printf("DEBUG MAIN could not read from server")
			os.Exit(1)
		}
		resultString = strings.TrimSuffix(resultString, "$")

		fmt.Printf("%s\n", resultString)

		for resultString != "End_Of_Connection" {
			resultString, err := reader.ReadString('$')
			if err != nil {
				fmt.Printf("DEBUG MAIN could not read from server")
				os.Exit(1)
			}
			resultString = strings.TrimSuffix(resultString, "$")

			fmt.Printf("%s\n", resultString)

		}

		//io.WriteString(conn, fmt.Sprintf("Coucou %d\n", i))  -> ecrire

		/*
					resultString, err := reader.ReadString('\n')
			            if (err != nil){
			                fmt.Printf("DEBUG MAIN could not read from server")
			                os.Exit(1)
			            }
			            resultString = strings.TrimSuffix(resultString, "\n")
		*/
		//=> lire

		/*
		    for i:= 0; i < 10; i++{

		        io.WriteString(conn, fmt.Sprintf("Coucou %d\n", i))

		        resultString, err := reader.ReadString('\n')
		        if (err != nil){
		            fmt.Printf("DEBUG MAIN could not read from server")
		            os.Exit(1)
		        }
		        resultString = strings.TrimSuffix(resultString, "\n")
		        fmt.Printf("#DEBUG server replied : |%s|\n", resultString)
		        time.Sleep(1000 * time.Millisecond)

		}*/

	}

}
