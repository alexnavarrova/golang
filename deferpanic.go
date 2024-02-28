package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Este es el mensaje final")
	fmt.Println("Este es el primer mensaje")
	a := 1
	if a == 1 {
		panic("hubo un error")
	} 
}

