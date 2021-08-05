package contas

import "github.com/lnl/banco/clientes"

// Estrutura da conta corrente do banco
type ContaCorrente struct {
	Titular clientes.Titular
	NumeroAgencia int
	NumeroConta int
	Saldo float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	if valor < 0 {
		return "Não é possível sacar um valor negativo"
	}

	podeSacar := valor <= c.Saldo

 	if podeSacar {
		c.Saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.Saldo += valor
		return "Deposito realizado com sucesso. Saldo:", c.Saldo
	} else {
		return "O valor do deposito não pode ser menor que zero. Saldo:", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor <= c.Saldo && valor > 0 {
		c.Saldo -= valor
		contaDestino.Depositar(valor)
		return true
	} else {
		return false
	}
}