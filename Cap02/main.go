package main

import "fmt"

func tarea(id int, canal chan<- string) {
	canal <- fmt.Sprintf("tarea ejecutada id:%d\n", id)
}
func main() {

	canal := make(chan string)

	//solo puedo enviar string por este canal

	for i := range 3 {
		go tarea(i, canal)
	}
	for range 3 {
		msg := <-canal
		fmt.Println(msg)
	}
}
