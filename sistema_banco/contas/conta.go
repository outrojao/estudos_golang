package contas

import (
	"fmt"
	"sistema_banco/clientes"
)

type Conta struct {
	Titular     *clientes.Titular
	numeroConta string
	saldo       float64
	transacoes  []string
}

func NewConta(titular *clientes.Titular, numeroConta string) (c *Conta) {
	c = &Conta{
		Titular:     titular,
		numeroConta: numeroConta,
		saldo:       0.0,
		transacoes:  []string{},
	}
	return
}

func (c *Conta) GetSaldo() float64 {
	return c.saldo
}

func (c *Conta) GetNumeroConta() string {
	return c.numeroConta
}

func (c *Conta) Depositar(valor float64) {
	c.saldo += valor
	c.transacoes = append(
		c.transacoes,
		fmt.Sprintf("Depósito: R$ %.2f", valor),
	)
}

func (c *Conta) Extrato() {
	fmt.Printf("Extrato da conta %s - Titular: %s\n", c.numeroConta, c.Titular.Nome)
	for _, transacao := range c.transacoes {
		fmt.Println(transacao)
	}
}

func (c *Conta) Sacar(valor float64) bool {
	if valor > c.saldo {
		return false
	}
	c.saldo -= valor
	c.transacoes = append(
		c.transacoes,
		fmt.Sprintf("Saque: R$ %.2f", valor),
	)
	return true
}

func (c *Conta) Transferir(valor float64, contaDestino *Conta) bool {
	if valor > c.saldo {
		return false
	}
	c.saldo -= valor
	contaDestino.saldo += valor
	c.transacoes = append(
		c.transacoes,
		fmt.Sprintf("Transferência: R$ %.2f para a conta %s", valor, contaDestino.numeroConta),
	)
	contaDestino.transacoes = append(
		contaDestino.transacoes,
		fmt.Sprintf("Recebimento de transferência: R$ %.2f da conta %s", valor, c.numeroConta),
	)
	return true
}
