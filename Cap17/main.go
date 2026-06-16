package main

import "fmt"

// Interfaces
type Procesador interface { //contrato interface
	Procesar(id int) string
}
type Impressor struct{} //tipo

func (i Impressor) Procesar(id int) string {
	return fmt.Sprintf("precesando un id %d", id)
}
func duplicar[T any](v T) []T {
	return []T{v, v}
}

//referentce vs value

func cambia(s []int) {
	s[0] = 99       //si afecta afuera
	s = []int{1, 2} //no afecta afuera
}

func main() {
	//Interfaces
	var p Procesador = Impressor{}
	fmt.Println(p.Procesar(42))

	//generics
	res := duplicar[string]("hola")
	fmt.Println(res)

	//refence vs values

	slice := []int{1, 2, 3}
	cambia(slice)
	fmt.Println("slice despues:", slice)
}
