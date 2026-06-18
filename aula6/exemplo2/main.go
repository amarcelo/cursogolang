package main

import "fmt"

func main() {

	// Slice vazio para armazenar nomes de alunos
	alunos := []string{}

	// Cadastrando alunos
	alunos = append(alunos, "Maria")
	alunos = append(alunos, "João")
	alunos = append(alunos, "Pedro")
	alunos = append(alunos, "José")

	fmt.Println("Lista de alunos:")

	// Percorre o slice exibindo os alunos
	for i, aluno := range alunos {
		fmt.Printf("%d - %s\n", i+1, aluno)
	}
}
