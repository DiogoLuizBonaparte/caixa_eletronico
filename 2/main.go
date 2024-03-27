package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calcularNotas(valor int) {
	if valor < 0 {
		fmt.Println("O valor inserido precisa ser maior que zero. Tente novamente")
		return
	}

	notas := []int{200, 100, 50, 20, 10, 5, 2}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	random := r.Perm(len(notas) - 1)

	for _, nota := range random {
		quantidade := valor / nota
		if quantidade > 0 {
			fmt.Printf("Notas de R$%d: %d\n", nota, quantidade)
			valor %= nota
		}
	}

	if valor != 0 {

		fmt.Printf("O valor restante de R$%d não pode ser sacado, pois não temos notas menor que R$2,00", valor)

	}
}

func main() {
	calcularNotas(273)
}
