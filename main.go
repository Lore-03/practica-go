package main

import "fmt"

func main() {

	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.15

	fmt.Println(pi, pi2)

	// Declaración de variables enteras

	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	//Area Cuadrada
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El area del cuadrado es:", areaCuadrado)

	//Operadores Aritmeticos
	x := 10
	y := 5

	//Suma
	resultado := x + y
	fmt.Println("El resultado de la suma es:", resultado)

	//Resta
	resultado = x - y
	fmt.Println("El resultado de la resta es:", resultado)

	//Multiplicacion
	resultado = x * y
	fmt.Println("El resultado de la multiplicacion es:", resultado)

	//Division
	resultado = x / y
	fmt.Println("El resultado de la division es:", resultado)

	//Modulo
	resultado = x % y
	fmt.Println("El resultado de el modulo es:", resultado)

	//Incrementable
	x++
	fmt.Println("El resultado de el incremento es:", x)

	//Decrementable
	y--
	fmt.Println("El resultado de el decremento es:", y)

	// reto 1
	//calcular el area de un rectangulo y de un circulo

	const baseRectangulo = 10
	const alturaRectangulo = 5
	const radioCirculo = 3

	areaRectangulo := baseRectangulo * alturaRectangulo
	areaCirculo := radioCirculo * radioCirculo

	fmt.Println("El area del rectangulo es:", areaRectangulo)
	fmt.Println("El area del circulo es:", areaCirculo)

	// Declaración de variables
	HelloMessage := "Hello"
	worldMessage := "World"
	cursos := 1000
	// Println imprime en una sola linea
	fmt.Println(HelloMessage, worldMessage)

	// Printf imprime en una sola linea pero con formato
	fmt.Printf("Hola %s %s\n", HelloMessage, worldMessage)

	//Sprintf se guarda en en la varaible message
	message := fmt.Sprintf("%s %s", HelloMessage, worldMessage)
	fmt.Println(message)

	//Tipo de dato
	fmt.Printf("HelloMessage: %T", HelloMessage)
	fmt.Printf("cursos: %T", cursos)
	//------------------------------

}
