package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	r := rand.NewSource(time.Now().UnixNano())
	s := rand.New(r)
	return s.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { //também suportado no switch
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu!!!")
	}
}
