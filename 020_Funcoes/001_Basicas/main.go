package main

import "fmt"

func minhaFuncao() {
	fmt.Println("F1: Esta é uma função sem retorno e sem parâmetro.")
}

func minhaFuncaoDois(parametroUm string, parametroDois string) {
	fmt.Printf("F2: Essa é uma função com parâmetro: %s %s \n", parametroUm, parametroDois)
}

func minhaFuncaoTres() string {
	return "F3: Essa é uma função com retorno."
}

func minhaFuncaoQuatro(parametroUm, parametroDois string) string {
	return fmt.Sprintf("F4: Essa função possui parâmetro e retorno: %s %s", parametroUm, parametroDois)
}

func minhaFuncaoCinco() (string, string) {
	return "F5: Função com dois retornos: Retorno 1.", "Retorno 2."
}

func main() {
	minhaFuncao()
	minhaFuncaoDois("Parâmetro 1.", "Parâmetro 2.")
	fmt.Println(minhaFuncaoTres())
	fmt.Println(minhaFuncaoQuatro("Parâmetro 1.", "Parâmetro 2."))
	retornoUm, retornoDois := minhaFuncaoCinco()
	fmt.Println(retornoUm, retornoDois)

}
