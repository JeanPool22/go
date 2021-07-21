package main

import "fmt"

func main() {
	/*for {
		fmt.Print("Salir? (s/n)")
		var c rune
		fmt.Scanf("%c\n", &c)
		if c == 'S' || c == 's' {
			break
		}
	}
	fmt.Println("Adiós!")

	for {
		fmt.Print("Salir? (s/n): ")
		var c rune
		fmt.Scanf("%c\n", &c)
		if c == 'N' || c == 'n' {
			continue
		}
		if c == 'S' || c == 's' {
			break
		}
		fmt.Println("carácter no reconocido")
	}
	fmt.Println("Adiós!")

	var c rune
	for c != 'S' && c != 's' {
		fmt.Print("Salir? (s/n): ")
		fmt.Scanf("%c\n", &c)
	}*/

	for i := 1; i <= 10 ; i++ {
		fmt.Println(i)
	}
}
