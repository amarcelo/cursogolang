package main

import "fmt"

// Estrutura que representa um aluno
type Aluno struct {
	Nome     string
	Idade    int
	Nota     float64
	Aprovado bool
}

func main() {

	// Cadastro dos alunos
	aluno1 := Aluno{
		Nome:     "Maria",
		Idade:    20,
		Nota:     8.5,
		Aprovado: true,
	}

	aluno2 := Aluno{
		Nome:     "João",
		Idade:    22,
		Nota:     4.0,
		Aprovado: false,
	}

	aluno3 := Aluno{
		Nome:     "Pedro",
		Idade:    19,
		Nota:     9.2,
		Aprovado: true,
	}

	// Exibindo os dados dos alunos

	fmt.Println("=== Aluno 1 ===")
	fmt.Println("Nome:", aluno1.Nome)
	fmt.Println("Idade:", aluno1.Idade)
	fmt.Println("Nota:", aluno1.Nota)
	fmt.Println("Aprovado:", aluno1.Aprovado)

	fmt.Println("\n=== Aluno 2 ===")
	fmt.Println("Nome:", aluno2.Nome)
	fmt.Println("Idade:", aluno2.Idade)
	fmt.Println("Nota:", aluno2.Nota)
	fmt.Println("Aprovado:", aluno2.Aprovado)

	fmt.Println("\n=== Aluno 3 ===")
	fmt.Println("Nome:", aluno3.Nome)
	fmt.Println("Idade:", aluno3.Idade)
	fmt.Println("Nota:", aluno3.Nota)
	fmt.Println("Aprovado:", aluno3.Aprovado)
}
