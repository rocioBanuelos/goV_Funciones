package main

import "fmt"

func main() {
	var cantImpares uint

	fmt.Println("Generador de números impares\n")
	fmt.Printf("Cantidad de números impares que desea: ")
	fmt.Scan(&cantImpares)

	if cantImpares > 0 {
		fmt.Println("\nLos números impares generados son: ")
		sigImpar := generarImpares()
		for i := 0; i < int(cantImpares); i++ {
			fmt.Println(sigImpar())
		}
	} else {
		fmt.Println("Ingrese un cantidad de números impares mayor a 0")
	}
}

func generarImpares() func() uint {
	i := uint(1) 
	return func() uint {
		var impar = i
		i += 2
		return impar
	}
}