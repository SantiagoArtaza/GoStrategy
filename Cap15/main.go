package main

import (
	"fmt"
	"unsafe"
)

//unsafe pointer

func main() {

	b := []byte("hola")
	s := *(*string)(unsafe.Pointer(&b))
	fmt.Println("", s)
}
