package main

import (
	"fmt"
	"sync"
)

// sync.wait froup espera a que todo termine
var (
	completadas int
	mu          sync.Mutex
)

func tarea(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	completadas++
	mu.Unlock()
	fmt.Sprintf("Tarea completadas :%d \n", completadas)
}

func main() {
	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go tarea(i, &wg)
	}
	wg.Wait()
	mu.Lock()
	fmt.Sprintf("Tarea completadas :%d \n", completadas)
	mu.Unlock()
}
