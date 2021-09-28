package main

import "runtime/debug"

func funcaoTres() {
	debug.PrintStack()
}

func funcaoDois() {
	funcaoTres()
}

func funcaoUm() {
	funcaoDois()
}

func main() {
	funcaoUm()
}
