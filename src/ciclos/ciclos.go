package main

import "fmt"

func main() {
	// ciclos for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}
	// for forever
	// counterForever := 0
	// for {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// }

	// ciclo for con iteración incrementable
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
}
