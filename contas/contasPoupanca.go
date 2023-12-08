package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (conta *ContaPoupanca) Sacar(valor float64) (string, float64) {
	if valor < conta.saldo && valor > 0 {
		conta.saldo -= valor
		return "Saque feito com sucesso", conta.saldo
	} else {
		return "Valor inválido", conta.saldo
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito menor que zero", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
