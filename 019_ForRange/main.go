package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5, 6} //compilador conta quantos elementos há dentro do array

	//aqui é acessado o índice e o valor de cada índice
	for i, numero := range numeros {
		fmt.Printf("Índice %d: %d \n", i, numero)
	}

	//aqui é acessado apenas o valor de cada índice sem precisar utilizar o índice em si
	for _, num := range numeros {
		fmt.Println(num)
	}
}
