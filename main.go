package main

import (
	"fmt"

	"github.com/lnl/banco/contas"
)

func main()  {
	contaLeandro := contas.ContaCorrente{
		Titular: "Leandro",
		NumeroConta: 123654,
		Saldo: 102.6,
	}

	contaLombi := contas.ContaCorrente {
		Titular: "Lombi",
		Saldo: 536.2,
	}

	fmt.Println(contaLeandro.Saldo)
	fmt.Println(contaLeandro.Sacar(30))
	fmt.Println(contaLeandro.Saldo)

	fmt.Println(contaLeandro.Depositar(50))

	fmt.Println(contaLombi)

	resultado := contaLeandro.Transferir(10, &contaLombi)
	fmt.Println(resultado)
	fmt.Println(contaLombi.Saldo)
}