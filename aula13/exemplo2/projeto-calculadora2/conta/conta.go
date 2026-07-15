// Package conta contém estruturas e funções
// relacionadas ao controle de contas bancárias.
package conta

import "errors"

// Conta representa uma conta bancária.
//
// O campo Titular começa com letra maiúscula.
// Portanto, pode ser acessado por outros pacotes.
//
// O campo saldo começa com letra minúscula.
// Portanto, só pode ser acessado dentro do pacote conta.
type Conta struct {
	Titular string
	saldo   float64
}

// NovaConta funciona como uma função construtora.
//
// Ela cria uma conta já com os dados iniciais.
// Como saldo é privado, seu valor é definido
// dentro do próprio pacote.
func NovaConta(titular string, saldoInicial float64) Conta {
	// Não permitimos saldo inicial negativo.
	if saldoInicial < 0 {
		saldoInicial = 0
	}

	return Conta{
		Titular: titular,
		saldo:   saldoInicial,
	}
}

// ConsultarSaldo devolve o saldo atual da conta.
//
// O saldo não é acessado diretamente.
// O acesso acontece por meio de um método público.
func (c Conta) ConsultarSaldo() float64 {
	return c.saldo
}

// Depositar acrescenta um valor ao saldo.
//
// Usamos um ponteiro *Conta porque o método
// precisa alterar a conta original.
func (c *Conta) Depositar(valor float64) error {
	if valor <= 0 {
		return errors.New("o valor do depósito deve ser maior que zero")
	}

	c.saldo += valor

	return nil
}

// Sacar retira um valor do saldo.
//
// O método valida o valor antes de alterar
// o estado interno da conta.
func (c *Conta) Sacar(valor float64) error {
	if valor <= 0 {
		return errors.New("o valor do saque deve ser maior que zero")
	}

	if valor > c.saldo {
		return errors.New("saldo insuficiente")
	}

	c.saldo -= valor

	return nil
}
