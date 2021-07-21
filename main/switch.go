package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("La arquitectura de su procesador es ")

	arch := runtime.GOARCH
	switch arch {
	case "386":
		fmt.Println("x86 de 32 bits")
	case "amd64":
		fmt.Println("x86 de 64 bits")
	default:
		fmt.Println(arch)
	}

	fmt.Println("Su sistema operativo es ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS X")
	case "linux":
		fmt.Println("GNU/Linux")
	case "hurd":
		fmt.Println("GNU/Hurd")
	default:
		fmt.Println(os)

	}
}
