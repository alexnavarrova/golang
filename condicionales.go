package main

import "fmt"

func main(){
	edad := 18
	if edad > 18 {
		fmt.Println("Eres mayor de 18")
	} else {
		fmt.Println("no eres mayor de edad")
	}


	// declarar variable  en una condicion
	if variable :=2 ; variable == 2 {
		fmt.Println("------------------------------------")
		fmt.Println("Variable es igual a 2")
		fmt.Println("variable se crea si cumple condicion, ", variable)
	}

	color := "redss"

	switch color {
		case "blue":
			fmt.Println("El color es azul")
		case "red":
			fmt.Println("El color es rojo")
		case "pink":
			fmt.Println("El color es rosado")
		case "green":
			fmt.Println("El color es verde")
		default:
			fmt.Println("El color es amarillo por defecto")
	}

	
}