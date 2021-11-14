package main

import (
	"fmt"
	"goProjetos/banco/contas"
)

func somarNumeros(valor string, num ...int) int {
	resultado := 0
	for _, numero := range num {
		resultado += numero
	}
	fmt.Printf("%s: %d\n", valor, resultado)
	return resultado
}

func PagarBoleto(conta VerificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type VerificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 1000)

	fmt.Println(contaDaLuisa.ObterSaldo())

}
