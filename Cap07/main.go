package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//sync map, y el atomic

func main() {
	var estado sync.Map //Loads, stores, and deletes run in amortized constant time.
	var completadas int64
	var wg sync.WaitGroup //HACE una cola de sucesos entonces con el wg.add agrega uno y hasta que no terminene todos no se termina el programa
	for i := range 3 {
		wg.Add(1) //AGREGO UNA MAS POR CADA UNO AGREGO UNO A LA ESPERA
		go func(id int) {
			defer wg.Done()
			estado.Store(fmt.Sprintf("tarea- %d", id), "completada") //guardanmos en el diccionario   tareria-1 : completada
			atomic.AddInt64(&completadas, 1)                         //(operacion atomica agrega una mas al completadas
			//atomic es una forma de hacer operaciones atomicas (sirve para sumar un contador sin logs)

		}(i)

	}
	wg.Wait()
	estado.Range(func(k, v any) bool {
		fmt.Println(k, v)
		return true
	})
	fmt.Printf("completadas %d ", completadas)

}
