package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	canal := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			canal <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}
	}()
	return canal
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			select {
			case s := <-entrada1:
				canal <- s
			case s := <-entrada2:
				canal <- s
			}
		}
	}()
	return canal
}

func main() {
	canal := juntar(falar("João"), falar("Maria"))
	fmt.Println(<-canal, <-canal)
	fmt.Println(<-canal, <-canal)
	fmt.Println(<-canal, <-canal)
}
