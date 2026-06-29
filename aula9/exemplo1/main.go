package main

import (
	"fmt"
	"time"
)

// Esta função representa o atendimento de um cliente.
// Ela recebe o ID do cliente e um channel para devolver o resultado.
func atenderCliente(id int, ch chan string) {

	// Simula um processamento demorado (consulta ao banco, API etc.)
	time.Sleep(2 * time.Second)

	// Envia a mensagem para a função main.
	ch <- fmt.Sprintf("Cliente %d atendido", id)
}

func main() {

	// Cria um channel para comunicação entre as goroutines e a main.
	ch := make(chan string)

	// Cria três goroutines, uma para cada cliente.
	for i := 1; i <= 3; i++ {
		go atenderCliente(i, ch)
	}

	// Aguarda os três atendimentos terminarem.
	for i := 1; i <= 3; i++ {
		fmt.Println(<-ch)
	}
}
