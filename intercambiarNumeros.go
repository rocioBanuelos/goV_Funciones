package main

import "fmt"

func main() {
	var a int64
	var b int64

	fmt.Println("Intercambiar dos números enteros\n")
	fmt.Printf("Ingrese el valor de 'a': ")
	fmt.Scan(&a)
	fmt.Printf("Ingrese el valor de 'b': ")
	fmt.Scan(&b)
	
	intercambia(&a, &b)

	fmt.Println("\nAhora el valor de 'a' es: ", a)
	fmt.Println("Ahora el valor de 'b' es: ", b)
}

func intercambia(a *int64, b *int64) {
	var temporal int64

	temporal = *a
	*a = *b
	*b = temporal
}