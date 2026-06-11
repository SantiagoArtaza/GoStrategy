package main

import "fmt"

type Tarea struct {
	ID          int
	Descripcion string
}
type TareaRepository interface {
	GuardarTarea(tarea Tarea) error
}

type MemRepo struct {
	tareas []Tarea
}

// aca tambien se puede crear una funcion para inciar el repositorio
func NewMemRepo() *MemRepo {
	return &MemRepo{
		tareas: []Tarea{},
	}
}

func (m *MemRepo) GuardarTarea(tarea Tarea) error {

	m.tareas = append(m.tareas, tarea)
	return nil
}

func ShowTask(repo TareaRepository) {
	if MemRepo, ok := repo.(*MemRepo); ok {
		for _, tarea := range MemRepo.tareas {
			fmt.Printf("%v\t%v\n", tarea.ID, tarea.Descripcion)
		}
	}
}
