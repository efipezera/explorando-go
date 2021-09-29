package main

import "fmt"

type item struct {
	idProduto  int
	quantidade int
	preco      float64
}

type pedido struct {
	idUsuario int
	itens     []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}
	return total
}

func main() {
	pedido := pedido{
		idUsuario: 1,
		itens: []item{
			{idProduto: 1, quantidade: 2, preco: 12.10},
			{idProduto: 2, quantidade: 1, preco: 23.49},
			{idProduto: 11, quantidade: 100, preco: 3.13},
		},
	}

	fmt.Printf("O valor total do pedido é de R$%.2f", pedido.valorTotal())
}
