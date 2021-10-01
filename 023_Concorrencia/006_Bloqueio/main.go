package main

import (
	"fmt"
	"time"
)

func rotina(canal chan int) {
	time.Sleep(time.Second)
	canal <- 1 //operação bloqueante
	fmt.Println("Só depois que foi lido.")
}

func main() {
	canal := make(chan int) //canal sem buffer

	go rotina(canal)

	fmt.Println(<-canal) //operação bloqueante
	fmt.Println("Foi lido.")
	fmt.Println(<-canal) //deadlock (acontece quando todas as goroutines estão mortas)
	fmt.Println("Fim.")  //nunca vai ser impresso
}
