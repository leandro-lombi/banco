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

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Deposito realizado com sucesso. Saldo:", c.saldo
	} else {
		return "O valor do deposito não pode ser menor que zero. Saldo:", c.saldo
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

func main()  {
	contaLeandro := ContaCorrente{
		titular: "Leandro",
		numeroConta: 123654,
		saldo: 102.6,
	}

	contaLombi := ContaCorrente {
		titular: "Lombi",
		saldo: 536.2,
	}

	fmt.Println(contaLeandro.saldo)
	fmt.Println(contaLeandro.Sacar(30))
	fmt.Println(contaLeandro.saldo)

	fmt.Println(contaLeandro.Depositar(50))

	fmt.Println(contaLombi)

	resultado := contaLeandro.Transferir(10, &contaLombi)
	fmt.Println(resultado)
	fmt.Println(contaLombi)
}