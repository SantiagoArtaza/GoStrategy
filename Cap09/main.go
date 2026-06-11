package main

import (
	"fmt"
	"log"
)

//defer,panic,recover,fatal

func puedeExplotar(divisor int) {
	defer func() {
		//recover te ayuda a proteger el panico y que no se rompa todo
		if err := recover(); err != nil {
			fmt.Println("Recuperando del panic:", err)
		}
	}()
	if divisor == 0 {
		//panic termina la ejecucion de la gorutine
		panic("divisor is zero")
	}
	fmt.Println("Resultado: ", 10/divisor)
}
func main() {
	puedeExplotar(2)
	puedeExplotar(0)
	log.Fatal("Esto termina el programa sin defer")

}
