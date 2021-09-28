package main

import "fmt"

func trocarValor(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return //retorno limpo
}

func main() {
	x, y := trocarValor(2, 3)
	fmt.Println(x, y)

	segundo, primeiro := trocarValor(7, 1)
	fmt.Println(segundo, primeiro)
}
