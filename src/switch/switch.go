package main

import "fmt"

func main() {
	modulo := 8 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//Sin condición
	value := 50
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor que 100")

	default:
		fmt.Println("Ninguna condicion")
	}

}
