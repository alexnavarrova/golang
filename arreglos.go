package main

import (
	"fmt"
	"reflect"
)

func main() {	
	fmt.Println("arreglos")

	var paises [4]string
	paises[0] = "Colombia"
	paises[1] = "Peru"
	paises[2] = "Venezuela"
	paises[3] = "Panama"
	fmt.Println(reflect.TypeOf(paises))
	fmt.Println(paises)
	fmt.Println(paises[2])
	fmt.Println("La cantidad",len(paises))

	fmt.Println("-----------slice-------------")
	var paises2 = make([]string, 5)
	paises2[0] = "Colombia"
	paises2[1] = "Peru"
	paises2[2] = "Venezuela"
	paises2[3] = "Ecuador"
	paises2[4] = "Bolivia"
	fmt.Println(paises2)

	// agregar elemento al slice
	paises2 = append(paises2, "Noruega")
	fmt.Println(paises2)

	//eliinar elementos
	//eliminar la posicion 5
	paises2 = append(paises2[:5], paises2[6:]...)
	fmt.Println(paises2)

	//eliminar la posicion 0
	paises2 = append(paises2[:0], paises2[1:]...)
	fmt.Println(paises2)

	//eliminar la posicion 2
	paises2 = append(paises2[:2], paises2[3:]...)
	fmt.Println(paises2)


	paises2 = append(paises2, "Noruega")
	paises2 = append(paises2, "Mexico")
	paises2 = append(paises2, "Guatemala")
	fmt.Println(paises2)

	paises2 = append(paises2[:0], paises2[3:]...)
	fmt.Println(paises2)

}