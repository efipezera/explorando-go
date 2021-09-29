package main

import "fmt"

//aqui a função faz uma cópia da variável
func incrementoUm(numero int) {
	numero++
	fmt.Println(numero)
	//return numero
}

func incrementoDois(numero *int) {
	/*
		* é usado para desreferenciar, ou seja, ter
		acesso ao valor no qual o ponteiro aponta.
		para ter acesso ao valor de numero,
		eu preciso utilizar o ponteiro.
	*/
	*numero++
}

func main() {
	n := 1

	//a função opera em cima de uma cópia do valor, fazendo que a variável n não seja afetada na chamada do incremento.
	incrementoUm(n) //por valor
	fmt.Println(n)

	//& é usado para obter o endereço da variável
	incrementoDois(&n) //por referência
	fmt.Println(n)
}
