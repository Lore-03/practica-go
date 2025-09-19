package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es palindromo")

	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {

	fmt.Println("Ingresar la palabra:")
	var palabra string
	//Scan leer datos desde la consola
	fmt.Scan(&palabra)
	//Tolower convierte todas las letras en mayusculas
	minu := strings.ToLower(palabra)
	isPalindromo(minu)
}
