package main

import (
	"fmt"

	"github.com/AlexDevauchelle/ELP_3_1/arrive"
)

func main() {
	fmt.Println("Hello world")
	var mavar = arrive.New([2]int{1, 2}, [2]int{3, 4}, float32(3.3), true)
	mavar.Affichage()

}
