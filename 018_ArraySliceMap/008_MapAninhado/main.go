package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 16468.58,
			"Gustavo Serra":  14986.98,
		},
		"H": {
			"Hugo Suzuki": 8670.90,
		},
		"I": {
			"Isabela Cabelo": 14290.98,
		},
	}

	delete(funcsPorLetra, "H")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
