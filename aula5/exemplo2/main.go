package main

import "fmt"

// Esta função recebe outra função como parâmetro
func executar(op func(int, int) int) {

	// Executa a função recebida usando os valores 10 e 5
	resultado := op(10, 5)

	fmt.Println("Resultado:", resultado)
}

// Função que soma dois números
func somar(a, b int) int {
	return a + b
}

func main() {

	// Estamos passando a função somar para executar
	executar(somar)

}
