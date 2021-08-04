package main

import "fmt"

// Estrutura da conta corrente do banco
type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float64
}

func main()  {
	// Inserir alguns campos específicos
	contaLeandro := ContaCorrente{
		titular: "Leandro",
		numeroConta: 123654,
		saldo: 102.6,
	}

	// Inserir todos os campos
	contaLeandro2 := ContaCorrente{"Lombi", 526, 654321, 6589.9}

	fmt.Println(contaLeandro)
	fmt.Println(contaLeandro2)
}