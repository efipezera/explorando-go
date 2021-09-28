package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	x := 10
	y := 20
	fmt.Printf("O resultado de %d + %d = %d", x, y, somar(x, y))
}
