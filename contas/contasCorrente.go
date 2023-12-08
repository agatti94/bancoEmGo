package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (conta *ContaCorrente) Depositar(valor float64) {
	conta.saldo = conta.saldo + valor
}

func (conta *ContaCorrente) ObterSaldo() float64 {
	return conta.saldo
}

func (conta *ContaCorrente) Sacar(valor float64) (string, float64) {
	if valor < conta.saldo && valor > 0 {
		conta.saldo -= valor
		return "Saque feito com sucesso", conta.saldo
	} else {
		return "Valor inválido", conta.saldo
	}
}

func (conta *ContaCorrente) Transeferir(valor float64, contaTransferida *ContaCorrente) (string, float64) {
	if valor <= conta.saldo && valor > 0 {
		conta.saldo -= valor
		contaTransferida.saldo += valor
		return "Transferência feita com sucesso", contaTransferida.saldo
	} else {
		return "Valor inválido", contaTransferida.saldo
	}
}
