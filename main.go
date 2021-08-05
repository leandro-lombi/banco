package main

import "fmt"

// Estrutura da conta corrente do banco
type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float64
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
		return "Saldo insuficiente"
	}
}

func main()  {
	contaLeandro := ContaCorrente{
		titular: "Leandro",
		numeroConta: 123654,
		saldo: 102.6,
	}

	fmt.Println(contaLeandro.saldo)
	fmt.Println(contaLeandro.Sacar(30))
	fmt.Println(contaLeandro.saldo)
}