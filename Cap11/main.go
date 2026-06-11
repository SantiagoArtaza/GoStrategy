package main

import (
	"context"
	"fmt"
	"time"
)

// context maneja  todas las cacelaciones y la deadlines
func tarea(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):

		fmt.Println("tarea completada")
	//ctx.Done() CANAL QUE SE CIERRA CUANDO EL CONTEXTO TERMINA
	case <-ctx.Done():
		fmt.Println("tarea cancelada", ctx.Err())
	}

}
func main() {
	//context.WithTimeout SE REFIERE A QUE VA A CANCELAR EL CONTEXTO DESPUES DE UN CIERTO TIEMPO
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	go tarea(ctx)
	time.Sleep(5 * time.Second)
}
