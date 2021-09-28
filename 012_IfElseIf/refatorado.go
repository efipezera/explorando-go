package main

import "fmt"

func notaParaConceito(nota float64) string {
	notaArredondada := int(nota)
	switch {
	case notaArredondada == 10 || notaArredondada == 9:
		return "Conceito A"
	case notaArredondada == 8 || notaArredondada == 7:
		return "Conceito B"
	case notaArredondada == 6 || notaArredondada == 5:
		return "Conceito C"
	case notaArredondada == 4 || notaArredondada == 3:
		return "Conceito D"
	case notaArredondada == 2 || notaArredondada == 1:
		return "Conceito E"
	default:
		return "Nota inválida!"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.9))
}
