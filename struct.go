package main

import (
	"fmt"
)

func main() {
	fmt.Println("--------------------struct-----------------------")
	estructura :=  Person{
		Id: 1,
		Nombre: "Cesar",
		Correo: "info@gmail.com",
		Edad: 42,
	}
	fmt.Println(estructura)
	fmt.Printf("%+v \n", estructura)
	fmt.Println(estructura.Nombre)

	p := new(Person)
	p.Id = 2
	p.Nombre = "Alex"
	p.Correo = "alexnavarrova@gmail.com"
	p.Edad = 15
	fmt.Println(p)
	fmt.Printf("%+v \n", p)
	fmt.Println(p.Correo)

}

//Inicia mayuscul para indicar que es publico
type Person struct {
	Id int
	Nombre string
	Correo string
	Edad int
}
