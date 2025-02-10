package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "ping"
	}
}

func pongar(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "pong"
	}
}


func imprimir(c chan string){
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep((time.Second * 1))
	}
}

func main() {
	var c chan string = make(chan string)

	go pingar(c)
	go imprimir(c)
	go pongar(c)

	var entrada string
	fmt.Scanln(&entrada)
}
