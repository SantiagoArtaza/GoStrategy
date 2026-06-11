package main

import (
	"fmt"
	"runtime"
	"time"
)

func tarea() {
	time.Sleep(100 * time.Microsecond)
}
func main() {

	fmt.Println("CPUs disponibles", runtime.NumCPU())
	fmt.Println("Goroutines antes", runtime.NumGoroutine())

	for i := range 10 {
		fmt.Println("go rutines", i)
		go tarea()
	}

	fmt.Println("Goroutines despues", runtime.NumGoroutine())
	runtime.Gosched() //director de orquesta , que dirige todas las gorutines usando los hilos del sistema
	time.Sleep(200 * time.Microsecond)
}
