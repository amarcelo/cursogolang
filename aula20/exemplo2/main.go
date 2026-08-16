package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	// Tentamos abrir o arquivo alunos.txt
	_, err := os.Open("alunos.txt")

	// Se aconteceu algum erro
	if err != nil {

		fmt.Println("Ocorreu um erro!")

		// ==========================================
		// errors.Is()
		// ==========================================
		// Pergunta:
		// "O erro aconteceu porque o arquivo
		// não existe?"

		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("IS: O arquivo não existe!")
		}

		// ==========================================
		// errors.As()
		// ==========================================
		// Criamos uma variável que poderá receber
		// um erro do tipo *os.PathError.

		var erroArquivo *os.PathError

		// Pergunta:
		// "O erro é do tipo *os.PathError?"
		//
		// Se for, coloque o erro dentro da
		// variável erroArquivo.

		if errors.As(err, &erroArquivo) {

			fmt.Println("\nAS: Detalhes do erro:")

			fmt.Println("Operação:", erroArquivo.Op)
			fmt.Println("Arquivo:", erroArquivo.Path)
			fmt.Println("Problema:", erroArquivo.Err)
		}
	}
}
