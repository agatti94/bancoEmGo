package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	var contaDoArthur *contas.ContaCorrente = new(contas.ContaCorrente)
	contaDoArthur.Titular.Nome = "Arthur"
	contaDoArthur.NumeroAgencia = 2
	contaDoArthur.NumeroConta = 3

	var contaDoJoca *contas.ContaCorrente = new(contas.ContaCorrente)
	contaDoArthur.Titular.Nome = "Joca"
	contaDoJoca.NumeroAgencia = 2
	contaDoJoca.NumeroConta = 3
	contaDoJoca.Saldo = 300
	contaDoArthur.Depositar(200)
	fmt.Println(contaDoArthur.Transeferir(100, contaDoJoca))
}
