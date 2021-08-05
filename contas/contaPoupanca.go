package contas

import "github.com/lnl/banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valor float64) string {
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

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Deposito realizado com sucesso. saldo:", c.saldo
	} else {
		return "O valor do deposito não pode ser menor que zero. saldo:", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
