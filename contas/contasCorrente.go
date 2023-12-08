package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (conta *ContaCorrente) Depositar(valor float64) {
	conta.Saldo = conta.Saldo + valor
}

func (conta *ContaCorrente) Sacar(valor float64) (string, float64) {
	if valor < conta.Saldo && valor > 0 {
		conta.Saldo -= valor
		return "Saque feito com sucesso", conta.Saldo
	} else {
		return "Valor inválido", conta.Saldo
	}
}

func (conta *ContaCorrente) Transeferir(valor float64, contaTransferida *ContaCorrente) (string, float64) {
	if valor <= conta.Saldo && valor > 0 {
		conta.Saldo -= valor
		contaTransferida.Saldo += valor
		return "Transferência feita com sucesso", contaTransferida.Saldo
	} else {
		return "Valor inválido", contaTransferida.Saldo
	}
}
