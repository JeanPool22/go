package main

import "fmt"

func main() {
	/*a := 3
	if a == 3 {
		i := 1
		fmt.Println("'i' solo es visible en este contexto:", i)
	}
	fmt.Println("'a' es visible en todo el 'main':", a)*/

	a := 0
	b := 0

	if true {
		a := 1
		b = 1
		a++
		b++
	}
	fmt.Printf("a = %d, b =%d\n", a,b)
}