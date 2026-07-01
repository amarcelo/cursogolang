package main

import (
	"fmt"
	"sync"
	"time"
)

// Quantidade total de pedidos
var totalPedidos int

// Mutex protege o contador
var mu sync.Mutex

func processarPedido(cliente int, wg *sync.WaitGroup) {

	// Informa que terminou
	defer wg.Done()

	fmt.Println("Cliente", cliente, "fazendo pedido...")

	// Simula processamento
	time.Sleep(500 * time.Millisecond)

	// Apenas uma goroutine pode alterar o contador
	mu.Lock()
	defer mu.Unlock()

	totalPedidos++

	fmt.Println("Pedido do cliente", cliente, "processado.")
}

func main() {

	var wg sync.WaitGroup

	// Simula 10 clientes acessando ao mesmo tempo
	for i := 1; i <= 10; i++ {

		wg.Add(1)

		go processarPedido(i, &wg)
	}

	// Espera todos terminarem
	wg.Wait()

	fmt.Println()
	fmt.Println("Total de pedidos:", totalPedidos)
}
