package main

import (
	"fmt"
	"time"
)

func isPrimo(numero int) bool {
	for i := 2; i < numero; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}

func primos(numeros int, canal chan int) {
	inicio := 2
	for i := 0; i < numeros; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				canal <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(canal)
}

func main() {
	canal := make(chan int, 30)
	go primos(cap(canal), canal)
	for primo := range canal {
		fmt.Printf("%d \n", primo)
	}
	fmt.Println("Fim.")
}
