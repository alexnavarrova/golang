package main

import "fmt"

func main() {
	fmt.Println("-------------maps--------")
	//{"id": 1, "nombre": "Colombia"}
	paises := make(map[string]int64)
	paises["Colombia"] = 50000000
	paises["Chile"] = 35000000
	paises["Venezuela"] = 45000000
	paises["España"] = 340000

	fmt.Println(paises)

	paises2 := map[int]string{
		1: "Colombia",
		2: "España",
		3: "Peru",
		4: "Bolivia",
		5: "Venezuela",
		11: "usa",
	}
	fmt.Println(paises2)

	pais, existe := paises2[11]

	if existe {
		fmt.Println("si existe el pais ", pais)
	} else {
		fmt.Println("no existe el pais ")
	}

	delete(paises2, 1)
	fmt.Println(paises2)

	// recorrer un map con el ciclo for
	for id, pais  := range paises2 {
		fmt.Println(id, pais)
	}

}