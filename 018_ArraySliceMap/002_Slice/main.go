package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	//slice NÃO é um array! slice define um pedaço de um array.

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]
	//esse é um slice do array a2, pegando no índice 1 até o índice 3, ou seja, os elementos: 2, 3
	fmt.Println(a2, s2)

	s3 := a2[:2]
	fmt.Println(a2, s3)

	//pode-se imaginar um slice como: tamanho e um ponteiro para um elemento de um array

	s4 := s3[:1]
	fmt.Println(s3, s4)

	//tudo isso são visualizações diferentes do mesmo array, no caso a2
}
