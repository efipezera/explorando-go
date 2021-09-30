package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	separacao := strings.Split(nomeCompleto, " ")
	p.nome = separacao[0]
	p.sobrenome = separacao[1]
}

func main() {
	p1 := pessoa{"Gisele", "Arruda"}
	fmt.Println(p1.getNomeCompleto())
	p1.setNomeCompleto("Maria Bezerra")
	fmt.Println(p1.getNomeCompleto())
}
