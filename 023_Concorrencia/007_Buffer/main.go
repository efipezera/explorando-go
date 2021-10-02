package main

import "fmt"

func rotina(canal chan int) {
	fmt.Println("Executou!")
	canal <- 1
	canal <- 2
	canal <- 3
	canal <- 4
	canal <- 5
	canal <- 6
}

func main() {
	//buffer serve para quando quiser enviar um determinado número de dados para o canal sem precisar ler cada um deles imediatamente.
	canal := make(chan int, 3)
	go rotina(canal)
	fmt.Println(<-canal)
}
