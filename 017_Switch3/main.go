package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Flutuante"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Não sei"
	}
}

func main() {
	fmt.Println(tipo(10))
	fmt.Println(tipo(2.5))
	fmt.Println(tipo("Olá"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
