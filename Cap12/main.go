package main

import (
	"bytes"
	"fmt"
)

// manipulaar ddatos rapidos slices internal, bites buffers
func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	//El arreglo y el slice esta sincronizados ya que uno nace del otro

	fmt.Println("slice:", slice)

	slice[0] = 9
	fmt.Println("arr despues de modificar el slice:", arr)

	//bytes.buffer
	var b bytes.Buffer //bufer lo que permite es manipular texto binario en memoria , de manera eficiente
	b.WriteString("Hello ")
	b.WriteString("world ")
	fmt.Println("buffer:", b.String())

}
