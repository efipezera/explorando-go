package main

import "fmt"

func main() {
	//atribuição longa
	var a byte = 3
	fmt.Println(a)

	//atribuição marmota
	b := 3 //inferência de tipo
	b += 3 //aditiva
	b -= 3 //subtrativa
	b *= 3 //multiplicativa
	b /= 3 //divisiva
	b %= 3 //modular
	fmt.Println(b)

	//atribuição de diversas variáveis
	x, y := 1, 2
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)

}
