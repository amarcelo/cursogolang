package main

import "fmt"

// Struct de endereço
type Endereco struct {
	Rua    string
	Numero int
	Cidade string
}

// Struct de funcionário
type Funcionario struct {
	Nome  string
	Idade int

	// Composição:
	// Funcionario possui um Endereco
	Endereco
}

func main() {

	funcionario := Funcionario{
		Nome:  "Antonio",
		Idade: 45,
		Endereco: Endereco{
			Rua:    "Rua das Flores",
			Numero: 100,
			Cidade: "Rio de Janeiro",
		},
	}

	fmt.Println("Nome:", funcionario.Nome)
	fmt.Println("Idade:", funcionario.Idade)

	// Acessando os campos da struct interna
	fmt.Println("Rua:", funcionario.Rua)
	fmt.Println("Número:", funcionario.Numero)
	fmt.Println("Cidade:", funcionario.Cidade)
}
