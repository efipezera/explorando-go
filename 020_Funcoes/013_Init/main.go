package main

import "fmt"

//essa função é inicializada antes da main
func init() {
	fmt.Println("Função init...")
}

func main() {
	fmt.Println("Função main...")
}
