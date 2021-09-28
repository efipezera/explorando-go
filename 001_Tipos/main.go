package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é:", reflect.TypeOf(32000))

	fmt.Println("--------------------")
	//sem sinal (só positivos)... uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	fmt.Println("--------------------")
	//com sinal (positivos e negativos)... int8, int16, int32, int64
	numeroMaximo := math.MaxInt64
	fmt.Println("O valor máximo do int64 é", numeroMaximo)
	fmt.Println("O tipo de numeroMaximo é", reflect.TypeOf(numeroMaximo))

	fmt.Println("--------------------")
	//rune: representa um mapeamento da tabela Unicode (int32)
	var exemploUm rune = 'a'
	fmt.Println("O valor de 'a' na tabela Unicode é", exemploUm)
	fmt.Println("O tipo de rune é", reflect.TypeOf(exemploUm))

	fmt.Println("--------------------")
	//números reais (float32, float64)
	var numeroFlutuante float64 = 69.99
	fmt.Println("O tipo de numeroFlutuante é", reflect.TypeOf(numeroFlutuante))

	fmt.Println("--------------------")
	//boolean
	booleano := true
	fmt.Println("O valor de booleano é", booleano)
	fmt.Println("O tipo de booleano é", reflect.TypeOf(booleano))

	fmt.Println("--------------------")
	//string
	texto := "Sou uma string"
	fmt.Println("O valor de texto é", texto)
	fmt.Println("O tipo de texto é", reflect.TypeOf(texto))
	fmt.Println("O tamanho da string é", len(texto))

	fmt.Println("--------------------")
	//string com múltiplas linhas
	multiplasLinhas := `Olá,
	meu nome é
	Felipe.`
	fmt.Println(multiplasLinhas)
	fmt.Println("O tamanho dessa string é", len(multiplasLinhas))

	fmt.Println("--------------------")
	//char
	fmt.Println("Char NÃO EXISTE!")
}
