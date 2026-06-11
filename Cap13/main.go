package main

import (
	"fmt"
	"reflect"
)

// automatizar validaciones y cerealizaciones
type Usuario struct {
	//`validate:"required"` metadatos
	Nombre string `validate:"required" db:"nombre"`
	Edad   int    `validate:"min=18" db:"edad"`
}

func main() {
	user := Usuario{"luca", 20}

	v := reflect.ValueOf(user)
	t := reflect.TypeOf(user)

	for i := 0; i < v.NumField(); i++ {
		campo := t.Field(i)
		valor := v.Field(i).Interface()
		tag := campo.Tag.Get("validate")
		fmt.Printf("campo:%s, valor: %v , tag: %s", campo.Name, valor, tag)

	}

}
