package contas

import "sistema_banco/clientes"

type ContaPoupanca struct {
	*Conta
	numeroAgencia string
}

func NewContaPoupanca(titular *clientes.Titular, numeroConta string, numeroAgencia string) (cp *ContaPoupanca) {
	cp = &ContaPoupanca{
		Conta:         NewConta(titular, numeroConta),
		numeroAgencia: numeroAgencia,
	}
	return
}

func (cp *ContaPoupanca) GetNumeroAgencia() string {
	return cp.numeroAgencia
}
