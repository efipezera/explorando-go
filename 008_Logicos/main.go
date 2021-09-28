package main

import "fmt"

func compras(trabalhoUm, trabalhoDois bool) (bool, bool, bool) {
	comprarTv50 := trabalhoUm && trabalhoDois
	comprarTv32 := trabalhoUm != trabalhoDois //ou exclusivo
	comprarSorvete := trabalhoUm || trabalhoDois

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv 50: %t, Tv 32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
