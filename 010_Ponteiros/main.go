package main

import "fmt"

func main() {
	i := 1

	var ponteiro *int = nil

	ponteiro = &i //pegando o endereço da variável i e colocando no ponteiro
	*ponteiro++   //desreferenciando (pegando o valor)
	i++

	//go não tem aritmética de ponteiro
	//ponteiro++
	fmt.Println(ponteiro, *ponteiro, i, &i)
}
