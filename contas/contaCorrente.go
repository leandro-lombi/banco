package contas

import "github.com/lnl/banco/clientes"

// Estrutura da conta corrente do banco
type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	if valor < 0 {
		return "Não é possível sacar um valor negativo"
	}

	podeSacar := valor <= c.saldo

	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Deposito realizado com sucesso. saldo:", c.saldo
	} else {
		return "O valor do deposito não pode ser menor que zero. saldo:", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor <= c.saldo && valor > 0 {
		c.saldo -= valor
		contaDestino.Depositar(valor)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
