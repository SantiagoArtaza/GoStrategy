package main

import (
	"fmt"
	"sync"
	"time"
)

// sync condition, permite que varias go rutines esperen a que se cumpla una condicion,
// wait va a dormir la go rutine hasta que reciba una se;al o broadcast,
// signal despierta una go rutine que estaba esperando solo una
// broadcast  despierta a todaas las que estan esperando
func main() {
	buf := NewBuffer()
	wg := sync.WaitGroup{}
	wg.Add(1)

	//consumidor
	go func() {
		defer wg.Done()
		buf.mu.Lock() // lo bloqueo porque lo voy a usar

		for len(buf.items) == 0 {
			buf.cond.Wait() //Espera hasta que haya item en el buffer
		}

		item := buf.items[0]      // saco el item del buffer, que es el primero
		buf.items = buf.items[1:] //saco el primer elemento del arreglo, practicamente consumiendo
		buf.mu.Unlock()

		fmt.Printf("Consumido: %d\n", item)
	}()

	time.Sleep(500 * time.Millisecond) //espera para que el consumidor se bloquee
	buf.mu.Lock()
	buf.items = append(buf.items, 42)
	buf.cond.Signal() //notifica al consumidor que hay un prod disponible
	buf.mu.Unlock()
	wg.Wait()

}
