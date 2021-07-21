package main

import (
	"fmt"
	"math/rand"
)

func main() {
	valor := rand.Int()
	if valor % 2 == 0 {
		fmt.Println("El número",valor,"es par")
	} else {
		fmt.Println("El número",valor,"es impar")
	}
	fmt.Println("Adiós!")
}