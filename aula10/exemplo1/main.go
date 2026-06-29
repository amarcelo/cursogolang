package main

import "fmt"

func main() {
	// Cria um channel bufferizado com espaço para 2 valores
	ch := make(chan int, 3)

	// Como o channel tem buffer 2, estes envios não bloqueiam
	ch <- 10
	ch <- 20

	// Agora lemos os valores do channel
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
