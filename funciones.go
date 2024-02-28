package main

import (
	"fmt"
	"time"
)

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
	fmt.Println(gorutines("Cesar", 1))
	time.Sleep(time.Second * 1)
	fmt.Println(gorutines("Alex",1))

	miCanal1 := make(chan string)
	miCanal2 := make(chan string)
	miCanal3 := make(chan string)
	go func(){
		miCanal1 <- gorutines("Pedro", 4)
	}()

	go func(){
		miCanal2 <- gorutines("Juan", 5)
	}()

	go func(){
		miCanal3 <- gorutines("Jaime", 6)
	}()

	fmt.Println(<-miCanal1)
	fmt.Println(<-miCanal2)
	fmt.Println(<-miCanal3)
	fmt.Println("Continuar con la ejecucion")

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


//goruntines
func gorutines(parametro string, tiempo int) string {
	time.Sleep(time.Second * time.Duration(tiempo))
	return "Hola " + parametro
}