package main

import (
	"fmt"
	"time"
)

// Simula uma API mais rápida
func apiRapida(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Resposta da API 1"
}

// Simula uma API mais lenta
func apiLenta(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Resposta da API 2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Executa as duas consultas ao mesmo tempo
	go apiRapida(ch1)
	go apiLenta(ch2)

	// O select espera vários channels
	// e executa o primeiro que estiver pronto
	select {
	case msg := <-ch1:
		fmt.Println(msg)

	case msg := <-ch2:
		fmt.Println(msg)
	}
}
