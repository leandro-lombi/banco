package main

import (
	"fmt"

	"github.com/lnl/banco/contas"
)

func main() {
	contaLeandro := contas.ContaCorrente{}

	contaLeandro.Depositar(100)

	fmt.Println(contaLeandro.ObterSaldo())
}
