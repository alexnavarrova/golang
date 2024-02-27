package main

import "fmt"

func main() {
	fmt.Println("clicos en golang")

	i := 1

	//como el while
	for i < 11 {
		fmt.Println(i)
		i++
	}
	fmt.Println("----------------------------------------------------")

	for i2:= 1; i2 < 11; i2++ {
		fmt.Println(i2)
		if i2%2 == 0 {
			fmt.Println("es par")
		}
		//continue
		//break
	}
}