package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//interfaces são implementadas implicitamente!

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$%.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Felipinho", "Souza"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Notebook", 6240.45}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	//p2 não é imprimivel
	p2 := produto{"Mouse", 145.60}
	imprimir(p2)
}
