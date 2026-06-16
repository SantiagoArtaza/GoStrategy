package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(500 * time.Millisecond)
	slapsed := time.Since(start) //cuanto paso desde un incio el timer
	fmt.Println("tardo:", slapsed)

	timer := time.NewTimer(1 * time.Second) //dispara un evente luego de tanto tiempo
	fmt.Println("esperando timer....")
	<-timer.C
	fmt.Println("timer expiro")
}
