package main

import (
	"fmt"
	"time"
)

func tarea(id int, canal chan<- string) {
	time.Sleep(time.Duration(id) * 300 * time.Millisecond)
	canal <- fmt.Sprintf("tarea %d terminada", id)
}
func main() {

	canal := make(chan string, 2) //bufer de 2 l canal con buffer lols que nos permite es encolar mensajes no bloquea hasta que se llena y el select espera por varios canales  ala vez

	for i := range 4 {
		go tarea(i, canal)
	}

	timeout := time.After(1 * time.Second)
	for range 4 {
		select {

		case msg := <-canal:
			fmt.Println(msg)

		case <-timeout:
			fmt.Println("tiempo de espera agorado")
		}
	}
}
