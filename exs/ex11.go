// Escreva um programa Golang que pede para o usuário informar um número inteiro e informa se este número é um número perfeito.
// Um número perfeito é aquele cuja soma dos seus divisores, exceto ele próprio, é igual ao número. Por exemplo, o número 6 é perfeito, pois 1 + 2 + 3 = 6.

package main

import "fmt"

func ehPerfeito(num int) bool {
	var soma int
	for i := 1; i < num; i++ {
		if num%i == 0 {
			soma += i
		}
	}

	if soma == num {
		return true
	}
	return false
}

func main() {
	fmt.Print("Digite um número para verificar se ele é perfeito: ")
	var num int
	fmt.Scanln(&num)

	if ehPerfeito(num) {
		fmt.Printf("%d É PERFEITO!", num)
	} else {
		fmt.Printf("%d NÃO É PERFEITO!", num)
	}
}
