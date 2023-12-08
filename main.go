package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta VerificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type VerificarConta interface {
	Sacar(valor float64) (string, float64)
}

func main() {
	var contaDoArthur *contas.ContaCorrente = new(contas.ContaCorrente)
	contaDoArthur.Titular.Nome = "Arthur"
	contaDoArthur.NumeroAgencia = 2
	contaDoArthur.NumeroConta = 3
	contaDoArthur.Depositar(500)

	PagarBoleto(contaDoArthur, 300)
	fmt.Println(contaDoArthur.ObterSaldo())

}
