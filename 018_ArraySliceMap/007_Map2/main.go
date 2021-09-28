package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Joãozinho":  11325.45,
		"Mariazinha": 15456.22,
		"Fulaninha":  18985.50,
	}

	funcsESalarios["Felipinho"] = 1350.56
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
