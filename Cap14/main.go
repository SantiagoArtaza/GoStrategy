package main

import "fmt"

func main() {
	// repo := &MemRepo{} //Es necesario el puntero porque estamos creando una instancia del memrepo ambas sirven
	repo := NewMemRepo()
	tarea := Tarea{ID: 1, Descripcion: "aprender go"}
	repo.GuardarTarea(tarea)

	fmt.Println("tarea guardada indivudal :", tarea)
	// fmt.Println("tareas guardadas en total :", repo.tareas)
	ShowTask(repo)

}
