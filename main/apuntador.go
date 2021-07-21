package main

import "fmt"

func main() {
	/*var pi *int
	pi = nil
	if pi == nil {
		fmt.Print("¡No puedo hacer nada con este apuntador ")
		fmt.Println("porque no apunta a nada!")
	}

	i := 10
	p := &i

	fmt.Println(i)
	fmt.Println(p)

	i := 10
	p := &i
	a := *p
	*p = 21
	fmt.Print(i,p,a,*p,i)*/

	i := 0
	j := 0
	p1 := &i
	p2 := &j
	if p1 == p2 {
		fmt.Println("p1 y p2 apuntan a la misma dirección")
	} else {
		fmt.Println("p1 y p2 apuntan a direcciones distintas")
	}
	if *p1 == *p2 {
		fmt.Println("p1 y p2 apuntan a valores iguales")
	} else {
		fmt.Println("p1 y p2 apuntan a valores distintos")
	}
}
