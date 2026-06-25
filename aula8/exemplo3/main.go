package main

import "fmt"

// Struct que representa uma conta bancária
type ContaBancaria struct {
	Titular string
	Saldo   float64
}

// Método com receiver por VALOR.
// Apenas consulta os dados da conta.
// Como não altera nada, não precisamos de ponteiro.
func (c ContaBancaria) MostrarSaldo() {
	fmt.Println("Titular:", c.Titular)
	fmt.Printf("Saldo: R$ %.2f\n", c.Saldo)
}

// Método com receiver por PONTEIRO.
// Como ele altera o saldo da conta,
// precisamos trabalhar com o objeto original.
func (c *ContaBancaria) Depositar(valor float64) {
	c.Saldo += valor
	fmt.Printf("Depósito de R$ %.2f realizado.\n", valor)
}

// Outro método com receiver por ponteiro.
// Retira dinheiro da conta.
func (c *ContaBancaria) Sacar(valor float64) {

	// Verifica se existe saldo suficiente
	if valor > c.Saldo {
		fmt.Println("Saldo insuficiente!")
		return
	}

	c.Saldo -= valor
	fmt.Printf("Saque de R$ %.2f realizado.\n", valor)
}

func main() {

	// Criando uma conta bancária
	conta := ContaBancaria{
		Titular: "Antonio",
		Saldo:   1000,
	}

	// Exibe o saldo inicial
	conta.MostrarSaldo()

	fmt.Println()

	// Deposita dinheiro
	conta.Depositar(500)

	// Mostra o novo saldo
	conta.MostrarSaldo()

	fmt.Println()

	// Realiza um saque
	conta.Sacar(300)

	// Exibe novamente o saldo
	conta.MostrarSaldo()
}
