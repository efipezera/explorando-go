package main

import (
	"fmt"
	"modulo/gerador"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		//Encaminhando os dados de origem para o destino.
		destino <- <-origem
	}
}

//Juntar tem a tarefa de multiplexar - misturar mensagens em um só canal.
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	canal := make(chan string)
	go encaminhar(entrada1, canal)
	go encaminhar(entrada2, canal)
	return canal
}

func main() {
	canal := juntar(
		gerador.Titulo("https://pt.stackoverflow.com/", "https://github.com/"),
		gerador.Titulo("https://www.google.com.br/", "https://www.youtube.com/"),
	)
	fmt.Println(<-canal, "|", <-canal)
	fmt.Println(<-canal, "|", <-canal)
}
