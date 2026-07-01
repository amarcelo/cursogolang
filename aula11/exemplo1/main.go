package main

import (
	"fmt"
	"sync"
)

// Saldo da conta compartilhado por todas as goroutines
var saldo = 1000

// Mutex protege o saldo
var mu sync.Mutex

func sacar(valor int) {

	// Trava o acesso ao saldo
	mu.Lock()

	// Garante que o Unlock será executado
	defer mu.Unlock()

	if saldo >= valor {
		saldo -= valor
		fmt.Println("Saque de", valor, "realizado.")
	} else {
		fmt.Println("Saldo insuficiente.")
	}
}

func main() {

	// Dois clientes tentando sacar ao mesmo tempo
	go sacar(100)
	go sacar(500)
	go sacar(200)

	// Apenas para dar tempo das goroutines executarem
	var entrada string
	fmt.Scanln(&entrada)

	fmt.Println("Saldo final:", saldo)
}
