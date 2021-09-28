package main

import (
	"fmt"
	"time"
)

func main() {
	//no go não existe while
	i := 1
	for i <= 10 {
		fmt.Printf("For while: %d \n", i)
		i++
	}

	//for tradicional
	for j := 0; j <= 20; j += 2 {
		fmt.Printf("For: %d \n", j)
	}

	//loop infinito
	for {
		fmt.Println("Infinito...")
		time.Sleep(time.Second * 2)
	}
}
