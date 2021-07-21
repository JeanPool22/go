package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if valor := rand.Int(); valor % 2 == 0 {
		fmt.Println("El número",valor,"es par")
	} else {
		fmt.Println("El número",valor,"es impar")
	}
	fmt.Println("Adiós!")
}
