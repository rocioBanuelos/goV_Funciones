package main

import "fmt"

func main() {
	var numero int64
	s := make([]int64, 0, 10)

	fmt.Println("Algoritmo de Ordenamiento Burbuja\n")
	fmt.Println("Ingrese los elementos que desea ordenar: ")
	for i := 0; i < 10; i++ {
		fmt.Scanf("%d\n", &numero)
		s = append(s,numero)
	}

	Burbuja(s)
}

func Burbuja(s []int64) {
	var temporal int64
	
	fmt.Println("Elementos sin ordenar: ", s)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				temporal = s[j]
				s[j] = s[j+1]
				s[j+1] = temporal
			}
		}
	}
	fmt.Println("Elementos ordenados: ", s)
}