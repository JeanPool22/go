package main

import "fmt"

func main() { 
	cosa := "deposito"
	x := 36
	y := 84
	fmt.Printf("Coordenadas de %v: (%b, %o)\n", cosa, x, y)

	/*var edad int
	fmt.Print("Edad? ")
	fmt.Scan(&edad)
	fmt.Println("Tienes", edad, "años")*/

	var hora, minuto, segundo int
	fmt.Print("HH:MM:SS? ")
	fmt.Scanf("%d:%d:%d", &hora, &minuto, &segundo)
	fmt.Printf("%d horas, %d minutos, %d segundos", hora, minuto, segundo)
}
