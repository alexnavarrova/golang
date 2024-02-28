package main

import (
	"fmt"
)

func main() {
	miFuncion(0)

}

func miFuncion(valor int) {
	dato := valor+ 1
	fmt.Println(dato)
	if dato < 10 {
		miFuncion(dato)
	}

}
