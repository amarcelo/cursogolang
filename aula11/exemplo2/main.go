package main

import (
	"fmt"
	"sync"
	"time"
)

// Simula o download de um arquivo
func baixar(nome string, wg *sync.WaitGroup) {

	// Informa que terminou quando a função acabar
	defer wg.Done()

	fmt.Println("Baixando:", nome)

	time.Sleep(2 * time.Second)

	fmt.Println(nome, "concluído!")
}

func main() {

	var wg sync.WaitGroup

	arquivos := []string{
		"foto.jpg",
		"video.mp4",
		"documento.pdf",
		"planilha.xlsx",
		"backup.zip",
	}

	for _, arquivo := range arquivos {

		// Adiciona uma tarefa
		wg.Add(1)

		// Executa em paralelo
		go baixar(arquivo, &wg)
	}

	// Espera todos terminarem
	wg.Wait()

	fmt.Println("Todos os downloads terminaram.")
}
