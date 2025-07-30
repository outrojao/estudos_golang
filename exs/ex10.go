// Escreva um programa GoLang que solicita ao usuário um número inteiro e informa se o valor informado é um número binário, ou seja, composto apenas pelos dígitos 0 e 1.

package main

import (
	"fmt"
	"strconv"
)

func verificaBinario(num int) bool {
	stringNum := strconv.Itoa(num)
	for i := 0; i < len(stringNum); i++ {
		if stringNum[i] != '0' && stringNum[i] != '1' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print("Digite um número binário: ")
	var numBinario int
	fmt.Scanln(&numBinario)
	if verificaBinario(numBinario) {
		fmt.Printf("%d É BINÁRIO", numBinario)
	} else {
		fmt.Printf("%d NÃO É BINÁRIO", numBinario)
	}
}
