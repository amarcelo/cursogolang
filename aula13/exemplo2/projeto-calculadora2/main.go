package main

import (
	"fmt"

	"cursogo.com/projeto-calculadora2/calculadora"
	"cursogo.com/projeto-calculadora2/conta"
)

func main() {
	fmt.Println("=== Demonstração de pacotes em Go ===")

	demonstrarCalculadora()
	demonstrarEncapsulamento()
}

// demonstrarCalculadora utiliza o primeiro pacote do projeto.
func demonstrarCalculadora() {
	fmt.Println("\n--- Pacote calculadora ---")

	numero1 := 10.0
	numero2 := 5.0

	soma := calculadora.Somar(numero1, numero2)
	subtracao := calculadora.Subtrair(numero1, numero2)
	multiplicacao := calculadora.Multiplicar(numero1, numero2)

	fmt.Printf("Soma: %.2f\n", soma)
	fmt.Printf("Subtração: %.2f\n", subtracao)
	fmt.Printf("Multiplicação: %.2f\n", multiplicacao)

	divisao, err := calculadora.Dividir(numero1, numero2)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Printf("Divisão: %.2f\n", divisao)
}

// demonstrarEncapsulamento utiliza o pacote conta.
func demonstrarEncapsulamento() {
	fmt.Println("\n--- Pacote conta ---")

	// Criamos uma conta por meio da função construtora.
	minhaConta := conta.NovaConta("Carlos", 1000)

	// Titular é um campo público.
	fmt.Println("Titular:", minhaConta.Titular)

	// O saldo é consultado por um método público.
	fmt.Printf(
		"Saldo inicial: R$ %.2f\n",
		minhaConta.ConsultarSaldo(),
	)

	// Tentativa de depósito.
	err := minhaConta.Depositar(500)

	if err != nil {
		fmt.Println("Erro no depósito:", err)
	} else {
		fmt.Println("Depósito realizado com sucesso.")
	}

	fmt.Printf(
		"Saldo após o depósito: R$ %.2f\n",
		minhaConta.ConsultarSaldo(),
	)

	// Tentativa de saque.
	err = minhaConta.Sacar(300)

	if err != nil {
		fmt.Println("Erro no saque:", err)
	} else {
		fmt.Println("Saque realizado com sucesso.")
	}

	fmt.Printf(
		"Saldo final: R$ %.2f\n",
		minhaConta.ConsultarSaldo(),
	)
}
