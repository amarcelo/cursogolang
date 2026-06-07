package main

import "fmt"

func main() {

	var nome string = "Antonio"
	var idade int = 20
	var nivel int = 90
	var vip bool = true

	fmt.Println("=== Sistema de Entrada do Jogo ===")
	fmt.Println("Jogador:", nome)

	// Verifica idade mínima
	if idade < 16 {
		fmt.Println("Acesso negado: idade insuficiente.")
	} else {

		fmt.Println("Idade permitida.")

		// Verifica benefícios do jogador
		if vip {
			fmt.Println("Você possui acesso VIP.")
		} else {
			fmt.Println("Você é um jogador comum.")
		}

		// Verifica classificação do jogador
		if nivel >= 50 {
			fmt.Println("Classe: Lendário")
		} else if nivel >= 30 {
			fmt.Println("Classe: Veterano")
		} else if nivel >= 10 {
			fmt.Println("Classe: Aventureiro")
		} else {
			fmt.Println("Classe: Noob")
		}
	}

	fmt.Println("Fim da verificação.")
}
