package main

import "fmt"

// Cria um contador de atendimentos
func criarContadorAtendimentos() func() int {

	// Esta variável ficará guardada pelo closure
	totalAtendimentos := 0

	// Retorna uma função
	return func() int {

		// Incrementa o total de atendimentos
		totalAtendimentos++

		// Retorna o número atual de atendimentos
		return totalAtendimentos
	}
}

func main() {

	// Cria o contador
	proximoAtendimento := criarContadorAtendimentos()

	fmt.Println("Cliente atendido número:", proximoAtendimento())
	fmt.Println("Cliente atendido número:", proximoAtendimento())
	fmt.Println("Cliente atendido número:", proximoAtendimento())
	fmt.Println("Cliente atendido número:", proximoAtendimento())

}
