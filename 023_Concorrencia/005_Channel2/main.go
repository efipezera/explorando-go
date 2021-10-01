package main

import (
	"fmt"
	"time"
)

//channel é a forma de comunicação entre as goroutines
//channel é um tipo

func doisTresQuatroVezes(base int, canal chan int) {
	time.Sleep(time.Second)
	canal <- 2 * base //enviando dados para o canal

	time.Sleep(time.Second)
	canal <- 3 * base

	time.Sleep(time.Second * 3)
	canal <- 4 * base
}
func main() {
	canal := make(chan int)
	go doisTresQuatroVezes(2, canal)

	a, b := <-canal, <-canal
	fmt.Println(a, b)

	fmt.Println(<-canal)
}
