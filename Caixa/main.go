package main

import "fmt"

func valDigitado() int {
	var input int
	fmt.Printf("Digite o valor a ser sacado ")
	fmt.Scanln(&input)
	return input
}

func getNotas() [7]int {
	var notas [7]int
	notas[0] = 200
	notas[1] = 100
	notas[2] = 50
	notas[3] = 20
	notas[4] = 10
	notas[5] = 5
	notas[6] = 2
	return notas
}

func umaNota(nota, numNotas int) {
	fmt.Printf("%d nota de R$%d,00\n", numNotas, nota)
}

func variasNotas(nota, numNotas int) {
	fmt.Printf("%d notas de R$%d,00\n", numNotas, nota)
}

func saque(valorSaque int) {
	if valorSaque < 2 {
		fmt.Println("Valor indisponível para saque. O valor precisa ser a partir de R$2,00. Tente novamente.")
		return
	}

	fmt.Printf("O saque será em: \n")
	for i, nota := range getNotas() {

		var numNotas = valorSaque / getNotas()[i]

		if numNotas == 1 {
			umaNota(nota, numNotas)
			valorSaque %= nota
		} else if numNotas >= 2 {
			variasNotas(nota, numNotas)
			valorSaque %= nota
		}
	}
	if valorSaque != 0 {
		fmt.Printf("O saldo restante da sua conta no valor de R$%d,00 não poderá ser sacado no momento, pois, as notas disponíveis são de R$2,00 em diante.", valorSaque)
	}
}

func main() {
	saque(valDigitado())
}
