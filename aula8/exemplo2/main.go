package main

import "fmt"

// Criando uma struct que representa um aluno
type Aluno struct {
	Nome  string
	Idade int
}

// Método com receiver por valor.
// O receiver é "a", que representa o objeto Aluno que chamou o método.
func (a Aluno) Apresentar() {
	fmt.Println("Meu nome é", a.Nome)
	fmt.Println("Tenho", a.Idade, "anos")
}

// Outro método da mesma struct.
func (a Aluno) FazerAniversario() {
	fmt.Println(a.Nome, "fez aniversário!")
	fmt.Println("A idade dentro do método é", a.Idade+1)
}

// Método com receiver por ponteiro.
// Como recebe um ponteiro (*Aluno), ele altera o objeto original.
func (a *Aluno) AtualizarIdade(novaIdade int) {
	a.Idade = novaIdade
}

func main() {

	// Criando um objeto da struct Aluno
	aluno := Aluno{
		Nome:  "Antonio",
		Idade: 30,
	}

	// Chamando o método Apresentar()
	aluno.Apresentar()

	fmt.Println()

	// Chamando outro método
	aluno.FazerAniversario()

	fmt.Println()

	// Alterando a idade utilizando um receiver por ponteiro
	aluno.AtualizarIdade(31)

	// O objeto original foi modificado
	aluno.Apresentar()
}
