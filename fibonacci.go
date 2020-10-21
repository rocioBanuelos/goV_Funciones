package main

import "fmt"

func main() {
	var pos_n uint64
	fmt.Println("Sucesión de Fibonacci\n")
	fmt.Printf("Ingrese el número de la posición hasta la cual desea visualizar (n): ")
	fmt.Scan(&pos_n)
	if pos_n == 0{
		fmt.Println("Fibonacci(",pos_n,") = ", Fibonacci(uint64(pos_n)))
	}else{
		for i:= 0; i < int(pos_n)+1; i++{
			fmt.Println("Fibonacci(",i,") = ", Fibonacci(uint64(i)))
		}
	}
}

func Fibonacci(n uint64) uint64 {
	if n == 0 || n == 1 {
		return n
	}else {
	return Fibonacci(n - 1) + Fibonacci(n - 2)
	}
}