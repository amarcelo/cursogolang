package main

import (
	"fmt"
	"time"
)

// Simula o download de um arquivo.
func baixar(nome string, ch chan string) {

	// Informa que o download foi iniciado.
	fmt.Println("Baixando", nome, "...")

	// Simula o tempo necessário para realizar o download.
	time.Sleep(2 * time.Second)

	// Informa para a função main que o download terminou.
	ch <- nome + " concluído!"
}

func main() {

	// Cria um channel para receber o resultado dos downloads.
	ch := make(chan string)

	// Inicia três downloads simultaneamente.
	go baixar("arquivo1.pdf", ch)
	go baixar("arquivo2.pdf", ch)
	go baixar("arquivo3.pdf", ch)

	// Aguarda os três downloads terminarem.
	for i := 0; i < 3; i++ {

		// Recebe a mensagem enviada pela goroutine.
		resultado := <-ch

		// Exibe o resultado na tela.
		fmt.Println(resultado)
	}

	fmt.Println("Todos os downloads foram concluídos!")
}
