package main

import "fmt"

func Suma(a, b int) int {
	return a + b
}

func Multiplica(a, b int) int {
	return a * b
}

func main() {
	var generador func() int
	var generadorDos func() int

	generador = func() int {
		return 0
	}
	fmt.Println("generador de ceros:", generador())

	contador := 0
	generadorDos = func() int {
		contador++
		return contador
	}
	fmt.Println("generadorDos contador:", generadorDos(), generadorDos(), generadorDos())

	var operador func(int, int) int
	operador = Suma
	fmt.Println("suma =", operador(3, 4))
	operador = Multiplica
	fmt.Println("multiplica =", operador(3, 4))
}
