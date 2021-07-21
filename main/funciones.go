package main

import "fmt"

func Hola(nombre string, apellido string)  {
	fmt.Printf("¡Hola, %s %s!\n", nombre, apellido)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxMin retorna dos números:
// Primero, el mayor de los argumentos
// Segundo, el menor de los argumentos
func MaxMin(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}

// MaxMinDos retorna dos números:
// Primero, el mayor de los argumentos
// Segundo, el menor de los argumentos
func MaxMinDos(a, b int) (max int, min int) {
	if a > b {
		max = a
		min = b
	} else {
		min = a
		max = b
	}
	return
}

func Incrementa(d *int)  {
	*d++
}

func main() {
	Hola("Marta", "García")
	Hola("Juan", "Martínez")
	m := Max(3, 5)
	a := 10 + Max(m + 2, 6)
	fmt.Printf("m: %d y a: %d\n", m, a)
	max, min := MaxMin(34, 55)
	fmt.Printf("max: %d y min: %d\n", max, min)
	mx, mn := MaxMinDos(3, 5)
	fmt.Printf("max: %d y min: %d\n", mx, mn)
	m, _ = MaxMinDos(3, 6)
	fmt.Printf("max: %d\n", m)
	d := 3
	fmt.Println("a =", d)
	Incrementa(&d)
	fmt.Println("a =", d)
}
