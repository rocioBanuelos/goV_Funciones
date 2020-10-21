package main

import "fmt"

func main() {
	fmt.Println("Retornar el número más grande")
	fmt.Println("El número más grande es: ", numeroMayor(0, 22, 12, 5, 3, 32, 6))
}

func numeroMayor(args ...int64) int64 {
	mayor := args[0]
	for _, valor := range args {
		if valor > mayor {
			mayor = valor
		}
	}
	return mayor
}