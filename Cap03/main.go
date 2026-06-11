package main

import (
	"fmt"
	"sync"
)

// race condicion
var (
	completadas int
	mu          sync.Mutex
)

func tarea(id int, canal chan<- string) {
	mu.Lock()
	completadas++
	mu.Unlock()
	canal <- fmt.Sprintf("Tarea %d completada", id)
}

func main() {
	canal := make(chan string)

	for i := range 5 {
		go tarea(i, canal)

	}
	for range 5 {
		fmt.Println(<-canal)
	}

	mu.Lock()
	fmt.Sprintf("Numero de tareas completas: %d", completadas)
	mu.Unlock()

}
