package funciones

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripeArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

//---------------------------------------
// funciones para los algoritmos que se hicieron de operadores aritmeticos para calcular el area de un rectangulo y de un circulo

// const baseRectangulo = 10
// 	const alturaRectangulo = 5
// 	const radioCirculo = 3

// 	areaRectangulo := baseRectangulo * alturaRectangulo
// 	areaCirculo := radioCirculo * radioCirculo

func calcularAreaRectangulo(base, altura int) int {
	areaRectangulo := base * altura
	return areaRectangulo
}

func calcularAreaCirculo(radio int) float64 {
	const pi = 3.14
	areaCirculo := pi * float64(radio) * float64(radio)
	return areaCirculo
}

func main() {
	normalFunction("Hola Mundo")
	tripeArgument(1, 2, "Hola Mundo")

	value := returnValue(2)
	fmt.Println(value)

	value1, value2 := doubleReturn(2)
	fmt.Println("value1 y value2:", value1, value2)
	//------------------------------------------------------------------

	areaRect := calcularAreaRectangulo(10, 6)
	fmt.Println("El área del rectángulo es:", areaRect)

	areaCirc := calcularAreaCirculo(6)
	fmt.Println("El área del círculo es:", areaCirc)
}
