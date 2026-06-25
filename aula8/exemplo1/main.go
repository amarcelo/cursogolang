package main

import "fmt"

// Criando uma struct
type Pessoa struct {
	Nome string
}

// Criando um método para a struct Pessoa.
// O receiver é "p", que representa o objeto que chamou o método.
func (p Pessoa) Apresentar() {
	fmt.Println("Olá! Meu nome é", p.Nome)
}

func main() {

	// Criando um objeto da struct Pessoa
	pessoa := Pessoa{
		Nome: "Antonio",
	}

	// Chamando o método Apresentar()
	pessoa.Apresentar()
}
