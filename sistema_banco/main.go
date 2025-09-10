package main

import (
	"fmt"
	cli "sistema_banco/clientes"
	con "sistema_banco/contas"
)

type verificarConta interface {
	Sacar(valor float64) bool
}

func PagarBoleto(valor float64, conta verificarConta) {
	conta.Sacar(valor)
}

func main() {
	cliente1 := cli.NewTitular("João Silva", "123.456.789-00", "Engenheiro")
	cliente2 := cli.NewTitular("Maria Oliveira", "987.654.321-00", "Médica")
	cliente3 := cli.NewTitular("Carlos Souza", "111.222.333-44", "Professor")

	conta := con.NewConta(cliente1, "12345-6")
	novaConta := con.NewConta(cliente2, "65432-1")
	outraConta := con.NewContaPoupanca(cliente3, "11223-3", "Agência Central")

	conta.Depositar(1000.0)
	conta.Depositar(250.0)
	novaConta.Depositar(500.0)
	novaConta.Sacar(100.0)
	novaConta.Sacar(600.0)

	conta.Transferir(300.0, novaConta)
	conta.Transferir(800.0, novaConta)
	novaConta.Transferir(200.0, outraConta.Conta)
	PagarBoleto(50.0, conta)

	conta.Extrato()
	novaConta.Extrato()
	outraConta.Extrato()

	fmt.Println("Saldo da conta de João:", conta.GetSaldo())
	fmt.Println("Saldo da conta de Maria:", novaConta.GetSaldo())
	fmt.Println("Saldo da conta de Carlos:", outraConta.GetSaldo())
	fmt.Println(conta != novaConta)
}
