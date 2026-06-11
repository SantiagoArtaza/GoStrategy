package main

import (
	"fmt"
	"runtime"
)

// recolector de basura 3 color

func main() {
	fmt.Println("memoria antes del GC garbage coleccter")
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("heap;alloc:%d bites\n ", m.HeapAlloc) //heap aloc incrementa cuando se acumula basura en el basurero y decrece cuando se va la basura

	runtime.GC() //fuerzsa el gc

	fmt.Println("memoria despues de gc")
	runtime.ReadMemStats(&m)
	fmt.Printf("heap;alloc:%d bites\n ", m.HeapAlloc)

}
