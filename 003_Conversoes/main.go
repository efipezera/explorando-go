package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado... o número entre os parentêses não é um int, mas sim um unicode!
	fmt.Println("Teste " + string(97))

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	//string para int
	num, _ := strconv.Atoi("10")
	fmt.Println("Agora num é do tipo:", reflect.TypeOf(num))

	b, _ := strconv.ParseBool("0")
	if b {
		fmt.Println("Verdadeiro.")
	} else {
		fmt.Println("Falso.")
	}
}
