package main

import (
	"fmt"

	"cursogo.com/projeto-calculadora/calculadora"
)

func main() {
	var numero1 float64
	var numero2 float64
	var operacao int

	fmt.Println("=== Calculadora em Go ===")

	fmt.Print("Digite o primeiro número: ")
	_, err := fmt.Scan(&numero1)

	if err != nil {
		fmt.Println("Primeiro número inválido.")
		return
	}

	fmt.Print("Digite o segundo número: ")
	_, err = fmt.Scan(&numero2)

	if err != nil {
		fmt.Println("Segundo número inválido.")
		return
	}

	fmt.Println("\nEscolha uma operação:")
	fmt.Println("1 - Somar")
	fmt.Println("2 - Subtrair")
	fmt.Println("3 - Multiplicar")
	fmt.Println("4 - Dividir")
	fmt.Print("Opção: ")

	_, err = fmt.Scan(&operacao)

	if err != nil {
		fmt.Println("Opção inválida.")
		return
	}

	switch operacao {
	case 1:
		resultado := calculadora.Somar(numero1, numero2)
		fmt.Printf("Resultado: %.2f\n", resultado)

	case 2:
		resultado := calculadora.Subtrair(numero1, numero2)
		fmt.Printf("Resultado: %.2f\n", resultado)

	case 3:
		resultado := calculadora.Multiplicar(numero1, numero2)
		fmt.Printf("Resultado: %.2f\n", resultado)

	case 4:
		resultado, err := calculadora.Dividir(numero1, numero2)

		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		fmt.Printf("Resultado: %.2f\n", resultado)

	default:
		fmt.Println("Operação inexistente.")
	}
}
