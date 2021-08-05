package main

import (
	"fmt"

	"github.com/lnl/banco/contas"
)

func main() {
	contaLeandro := contas.ContaCorrente{}

	contaLeandro.Depositar(100)

	fmt.Println(contaLeandro.ObterSaldo())

	contaLombi := contas.ContaPoupanca{}

	fmt.Println(contaLombi)
	contaLombi.Depositar(50)
	fmt.Println(contaLombi.ObterSaldo())
	contaLombi.Sacar(15)
	fmt.Println(contaLombi.ObterSaldo())

}
