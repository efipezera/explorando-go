package main

import (
	"fmt"
	"time"
)

func main() {
	tempo := time.Now()
	switch { //switch true
	case tempo.Hour() < 12:
		fmt.Println("Bom dia!")
	case tempo.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite!")
	}
}
