package main

import (
	"fmt"
	"time"
)

// Simula o download de um arquivo.
func baixar(nome string, ch chan string) {

	// Simula o tempo necessário para baixar o arquivo.
	time.Sleep(2 * time.Second)

	// Informa que o download terminou.
	ch <- nome + " concluído"
}

func main() {

	// Cria o channel que receberá os resultados.
	ch := make(chan string)

	// Cada download acontece em uma goroutine diferente.
	go baixar("foto.jpg", ch)
	go baixar("video.mp4", ch)
	go baixar("documento.pdf", ch)

	// Aguarda todos os downloads terminarem.
	for i := 0; i < 3; i++ {

		// Recebe e imprime o resultado.
		fmt.Println(<-ch)
	}
}
