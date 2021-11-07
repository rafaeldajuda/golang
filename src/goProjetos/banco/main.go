package main

import (
	"fmt"
	//c "goProjetos/banco/contas"
	c "goProjetos/banco/contas"
)

func somarNumeros(valor string, num ...int) int {
	resultado := 0
	for _, numero := range num {
		resultado += numero
	}
	fmt.Printf("%s: %d\n", valor, resultado)
	return resultado
}

func main() {

	contaDaSilvia := c.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := c.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDoGustavo.Trasnferir(-200, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

}
