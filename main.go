package main

import (
	"fmt"

	"github.com/lnl/banco/contas"
)

func PagarBoleto(conta verificarConta, valor float64) {
	conta.Sacar(valor)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaLeandro := contas.ContaCorrente{}

	contaLeandro.Depositar(100)

	fmt.Println(contaLeandro.ObterSaldo())

	PagarBoleto(&contaLeandro, 82)

	fmt.Println(contaLeandro.ObterSaldo())

	contaLombi := contas.ContaPoupanca{}

	fmt.Println(contaLombi)
	contaLombi.Depositar(50)
	fmt.Println(contaLombi.ObterSaldo())
	contaLombi.Sacar(15)
	fmt.Println(contaLombi.ObterSaldo())

	PagarBoleto(&contaLombi, 8)

	fmt.Println(contaLombi.ObterSaldo())
}
