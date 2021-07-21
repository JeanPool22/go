package main

import "fmt"

func main() {
	fmt.Print("Escribe un carácter: ")
	var c int8
	fmt.Scanf("%c", &c)

	switch {
	case c >= 'A' && c <= 'Z':
		fmt.Println("Letra Mayúscula")
	case c >= 'a' && c <= 'z':
		fmt.Println("Letra Minúscula")
	case c >= '0' && c <= '9':
		fmt.Println("Dígito")
	default:
		fmt.Println("Ni letra ni dígito")
	}
}
