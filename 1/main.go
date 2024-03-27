package main

import (
	"fmt"
)

func valDigitado() int {
	var input int
	fmt.Printf("Digite o valor a ser sacado ")
	fmt.Scanln(&input)
	return input
}

func saque(valorSaque int) {
	if valorSaque < 2 {
		fmt.Println("O valor do saque precisa ser a partir de R$2,00. Tente novamente.")
		return
	}
	var notas [6]int
	notas[0] = 200
	notas[1] = 100
	notas[2] = 50
	notas[3] = 20
	notas[4] = 10
	notas[5] = 2

	fmt.Printf("O saque será em: \n")
	for i, nota := range notas {

		var numNotas = valorSaque / notas[i]

		if numNotas == 1 {
			fmt.Printf("%d nota de R$%d,00\n", numNotas, nota)
			valorSaque %= nota
		} else if numNotas > 0 && numNotas != 1 {
			fmt.Printf("%d notas de R$%d,00\n", numNotas, nota)
			valorSaque %= nota
		}
	}
	if valorSaque != 0 {
		fmt.Printf("O saldo atual da sua conta no valor de R$%d,00 não pode ser sacado, pois, as notas disponíveis são de R$2,00 em diante.", valorSaque)
	}
}

func main() {
	saque(valDigitado())
}
