package main

import "sync"

// sync.cond, Signal,Broadcast, wait
// sync siempre cosas de sincornizacion
type buffer struct {
	mu    sync.Mutex
	cond  *sync.Cond
	items []int
}

func NewBuffer() *buffer {
	b := &buffer{items: []int{}}
	b.cond = sync.NewCond(&b.mu) //esto seria, que se genere una condicion  mientras esto no este bloqueado
	return b
}
