// Escreva um programa Golang que usa o laço for para solicitar ao usuário que informe 10 números inteiros. Em seguida mostre o maior e o menor valor lido. Não é permitido usar vetores ou matrizes (arrays).

package main

import "fmt"

func main() {
	fmt.Println("A seguir, digite 10 números inteiros:")
	var maiorNum, menorNum int
	for i := 1; i <= 10; i++ {
		var numAtual int
		fmt.Printf("%d° número: ", i)
		fmt.Scanln(&numAtual)

		if numAtual > maiorNum {
			maiorNum = numAtual
		} else if numAtual < menorNum {
			menorNum = numAtual
		}
	}
	fmt.Printf("MAIOR NÚMERO - %d\n", maiorNum)
	fmt.Printf("MENOR NÚMERO - %d", menorNum)
}
