package main

import "fmt"

// Esta função representa a cozinha.
// Ela recebe pedidos através do channel.
func cozinha(pedidos chan string) {

	// Enquanto houver pedidos no channel...
	for pedido := range pedidos {

		// Prepara o pedido recebido.
		fmt.Println("Preparando:", pedido)
	}
}

func main() {

	// Cria um channel para enviar pedidos.
	pedidos := make(chan string)

	// A cozinha começa a funcionar em paralelo.
	go cozinha(pedidos)

	// O caixa envia os pedidos para a cozinha.
	pedidos <- "Pizza Calabresa"
	pedidos <- "Pizza Portuguesa"
	pedidos <- "Pizza Frango"
	pedidos <- "Pizza Marguerita"

	// Fecha o channel indicando que não haverá novos pedidos.
	close(pedidos)
}
