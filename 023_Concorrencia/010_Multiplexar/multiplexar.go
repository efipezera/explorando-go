package main

import (
	"fmt"

	"github.com/efipezera/generator"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

//multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	canal := make(chan string)
	go encaminhar(entrada1, canal)
	go encaminhar(entrada2, canal)
	return canal
}

func main() {
	canal := juntar(
		generator.Titulo("https://pt.stackoverflow.com/", "https://github.com/"),
		generator.Titulo("https://www.google.com.br/", "https://www.youtube.com/"),
	)
	fmt.Println(<-canal, "|", <-canal)
	fmt.Println(<-canal, "|", <-canal)
}
