package main

import "fmt"

//clojure que retorna uma função
func clojure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	//pelo fato de existir clojure no go, a função vai saber de suas origens e vai utilizar o valor de x dado na clojure e não nesse novo escopo main.
	imprimeX := clojure()
	imprimeX()
}
