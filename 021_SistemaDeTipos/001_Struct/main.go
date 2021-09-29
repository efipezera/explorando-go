package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	desconto := p.preco / 100 * p.desconto
	valorFinal := p.preco - desconto
	return valorFinal
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Computador",
		preco:    6890.99,
		desconto: 10,
	}
	fmt.Printf("%v, preço final com desconto de %.f%%: R$%.2f \n", produto1, produto1.desconto, produto1.precoComDesconto())

	produto2 := produto{"Teclado", 158.98, 20}
	fmt.Printf("%v, preço final com desconto de %.f%%: R$%.2f \n", produto2, produto2.desconto, produto2.precoComDesconto())
}
