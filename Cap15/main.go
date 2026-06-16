package main

import (
	"fmt"
	"unsafe"
)

//unsafe pointer

func main() {

	b := []byte("hola")
	s := *(*string)(unsafe.Pointer(&b))
	fmt.Println("string convertido con unsafe", s) //lo que hace e ussafe  esque obliga a convertir, lo cual puede ser miuyu peligroso

}
