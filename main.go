package main

import (
	"fmt"
	"reflect"
)

const MiConstante = "Sol"

func main() {
	//declaracion por inferencia
	var nombre string = "Alexander"
	textoGrande := `Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum Lorem ipsum`
	fmt.Println(textoGrande)

	var estado bool = false
	fmt.Println(estado)

	var flotante32 float32 = 32.33
	fmt.Println(flotante32)

	var flotante64 float32 = 32.33
	fmt.Println(flotante64)

	var entero_int8 int = 8
	fmt.Println(entero_int8)

	var entero_int16 int16 = 16
	fmt.Println(entero_int16)

	var entero_int32 int32 = 32
	fmt.Println(entero_int32)

	var entero_int64 int64 = 64
	fmt.Println(entero_int64)


	//reflect.TypeOf() obtener el tipo de dato
	fmt.Println("El tipo de dato es: ",reflect.TypeOf(entero_int64))

	//puntero
	fmt.Println(entero_int16, &entero_int16)



	// declaracion corta o rapida
	apellido := "Navarro"

	fmt.Println(nombre + apellido)
	fmt.Printf("El valor de mi constante es %s\n", MiConstante)
}