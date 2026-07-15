// Package calculadora contém funções para realizar
// operações matemáticas básicas.
package calculadora

import "errors"

// Somar recebe dois números e devolve o resultado da soma.
//
// Como o nome começa com letra maiúscula,
// a função pode ser utilizada por outros pacotes.
func Somar(a float64, b float64) float64 {
	return a + b
}

// Subtrair recebe dois números e devolve
// o resultado da subtração.
func Subtrair(a float64, b float64) float64 {
	return a - b
}

// Multiplicar recebe dois números e devolve
// o resultado da multiplicação.
func Multiplicar(a float64, b float64) float64 {
	return a * b
}

// Dividir recebe dois números.
//
// A função devolve dois valores:
// 1. O resultado da divisão.
// 2. Um possível erro.
//
// Caso o divisor seja zero, a função devolve um erro.
func Dividir(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("não é possível dividir por zero")
	}

	return a / b, nil
}

// validarNumero é uma função auxiliar.
//
// Como seu nome começa com letra minúscula,
// ela só pode ser utilizada dentro deste pacote.
func validarNumero(numero float64) bool {
	return numero >= 0
}
