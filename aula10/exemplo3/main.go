package main

import (
	"fmt"
	"time"
)

// Worker é quem processa as tarefas da fila
func worker(id int, tarefas chan int, resultados chan string) {

	// O range continua lendo enquanto o channel tarefas estiver aberto
	for tarefa := range tarefas {

		fmt.Println("Worker", id, "processando tarefa", tarefa)

		// Simula o tempo de processamento
		time.Sleep(2 * time.Second)

		// Envia o resultado para outro channel
		resultados <- fmt.Sprintf("Worker %d concluiu tarefa %d", id, tarefa)
	}
}

func main() {

	// Channel bufferizado com espaço para 5 tarefas
	tarefas := make(chan int, 5)

	// Channel para receber os resultados
	resultados := make(chan string)

	// Criamos 2 workers trabalhando em paralelo
	go worker(1, tarefas, resultados)
	go worker(2, tarefas, resultados)

	// Enviamos 5 tarefas para a fila
	for i := 1; i <= 5; i++ {
		tarefas <- i
	}

	// Fechamos o channel para avisar que não haverá mais tarefas
	close(tarefas)

	// Esperamos os 5 resultados
	for i := 1; i <= 5; i++ {

		// O select espera o resultado ou um timeout
		select {
		case msg := <-resultados:
			fmt.Println(msg)

		case <-time.After(3 * time.Second):
			fmt.Println("Timeout: uma tarefa demorou demais")
		}
	}

	fmt.Println("Programa finalizado")
}
