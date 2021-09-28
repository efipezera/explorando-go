package main

import "fmt"

func main() {
	x := 1
	y := 2

	//o go tem apenas a forma postfix, não existe ++x, só x++
	x++
	fmt.Println(x)
	y--
	fmt.Println(y)
}
