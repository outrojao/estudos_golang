package clientes

type Titular struct {
	Nome      string
	CPF       string
	Profissao string
}

func NewTitular(nome string, cpf string, profissao string) (t *Titular) {
	t = &Titular{
		Nome:      nome,
		CPF:       cpf,
		Profissao: profissao,
	}
	return
}
