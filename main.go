package main

import (
	"fmt"

	"github.com/lnl/banco/contas"
	"github.com/lnl/banco/clientes"
)

func main()  {
	contaLeandro := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome: "Leandro",
			CPF: "653.325.326-40",
			Profissao: "Analista",
		},
		NumeroConta: 123654,
		Saldo: 102.6,
	}

	clienteLombi := clientes.Titular{
		Nome: "Lombi",
		CPF: "123.456.789-40",
		Profissao: "Desenvolvedor",
	}

	contaLombi := contas.ContaCorrente{clienteLombi, 1236, 6989, 536.2}

	fmt.Println(contaLeandro.Saldo)
	fmt.Println(contaLeandro.Sacar(30))
	fmt.Println(contaLeandro.Saldo)

	fmt.Println(contaLeandro.Depositar(50))

	fmt.Println(contaLombi)

	resultado := contaLeandro.Transferir(10, &contaLombi)
	fmt.Println(resultado)
	fmt.Println(contaLombi.Saldo)
}