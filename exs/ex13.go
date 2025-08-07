/*
Escreva um programa Golang que leia duas notas (como float64), calcule e mostre a média aritmética e uma mensagem de acordo com as seguintes regras:

1) Se a média for inferior a 4,0 escreva "Reprovado";
2) Se a média for igual ou superior a 4,0 e inferior a 7,0 escreva "Exame";
3) Se a média for igual ou superior a 7,0 escreva "Aprovado".
*/

package main

import "fmt"

func calculaMedia(nota1, nota2 float64) float64 {
	return (nota1 + nota2) / 2
}

func analisaNota(media float64) string {
	if media < 4.0 {
		return "Reprovado"
	} else if media < 7.0 {
		return "Exame"
	} else {
		return "Aprovado"
	}
}

func main() {
	fmt.Println("Digite os valores das 1° e 2° notas abaixo:")
	fmt.Print("1° Nota: ")
	var nota1 float64
	fmt.Scanln(&nota1)

	fmt.Print("2° Nota: ")
	var nota2 float64
	fmt.Scanln(&nota2)

	fmt.Printf("SITUAÇÃO DO ALUNO: %s", analisaNota(calculaMedia(nota1, nota2)))
}
