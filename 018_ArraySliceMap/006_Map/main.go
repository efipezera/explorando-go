package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//maps devem ser inicializados!
	//map é chave e valor
	aprovados := make(map[int]string)
	aprovados[123456789] = "Joãozinho"
	aprovados[987654321] = "Fulaninha"
	aprovados[975312468] = "Mariazinha"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Println("-------------------------")
		fmt.Printf("%s, CPF: %d \n", nome, cpf)
	}
	fmt.Println("-------------------------")
	//para deletar uma chave:
	delete(aprovados, 123456789)
	fmt.Println(aprovados)
}
