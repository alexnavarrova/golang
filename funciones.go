package main

import "fmt"

func main() {
	miFuncion()
	sumar(1,4)
	nombre := miFuncionConRetorno("Alexander")
	fmt.Println(nombre)
	nombre, apellido, edad := miFuncionConRetornoMultiple()

	fmt.Println(nombre, apellido, edad)
	fmt.Println("La suma es ", suma(22,56))
	fmt.Println("--------------------------------------------------------------")
	Table := tabla(2)
	for i := 1; i <= 10; i++ {
		fmt.Printf("2 x %v = %v ", i, Table())
		fmt.Println("")
	}

}

func miFuncion() {
	fmt.Println("desde mi funcion")
}

func sumar(n1 int, n2 int) {
	resultado := n1 + n2
	fmt.Println("La suma es ", resultado)
}

func miFuncionConRetorno(nombre string) string{
	return "Tu nombre es " + nombre
}

func miFuncionConRetornoMultiple() (string, string, int){
	return "Alex", "Navarro", 15
}

//funcion anonima
var suma = func(n1 int, n2 int) int {
	return n1 + n2
}

//clousure - funciones que retornan otra funcion
func tabla(n1 int) func() int {
	numero := n1
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
