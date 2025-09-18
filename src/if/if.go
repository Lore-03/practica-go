package main

import "fmt"

//Hacer una función que determine que un numero es par o impor /// acceso de usuario y contraseña

// Función para numeros pares e imparas

func numerosPares(number int) bool {
	return number%2 == 0
}

func platForm(Username, Password string) bool {
	return Username == "Vivalavida" && Password == "Dinosaurio123"

}

func main() {

	if numerosPares(8) {
		fmt.Println("Es par")
	} else {
		fmt.Println("Es impar")
	}
	//-----------------------------
	if platForm("Vivalavida", "Dinosaurio123") {
		fmt.Println("Usuario Correcto")
	} else {
		fmt.Println("Usuario incorrecto")
	}

	// valor1 := 1
	// valor2 := 2

	// if valor1 == 8 {
	// 	fmt.Println("es 1")
	// } else {
	// 	fmt.Println("No es 1")
	// }
	// //With and
	// if valor1 == 1 && valor2 == 2 {
	// 	fmt.Println("Es verdad")
	// }
	// // With Or
	// if valor1 == 0 || valor2 == 2 {
	// 	fmt.Println("Es verdad el o")
	// }

	// //Convertir texto a numero
	// value, err := strconv.Atoi("53")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Value", value)

}
