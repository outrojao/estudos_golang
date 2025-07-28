// Escrever um programa em Go que leia o nome de um aluno e as notas das três provas que ele obteve no semestre. No final informar o nome do aluno e a sua média (ponderada), onde o primeiro e o segundo teste têm pesos de 30% e o terceiro 40%.

package main
import (
  "fmt"
  "bufio"
  "os"
)

func calcMediaPonderada(nota1, nota2, nota3 float64) (media float64){
	nota1 = (30/100.0) * nota1
	nota2 = (30/100.0) * nota2
	nota3 = (40/100.0) * nota3

	media = nota1 + nota2 + nota3
	return
}

func main(){
	var nota1, nota2, nota3 float64
	var nomeAluno string

	// para ler a entrada do nome do usuário
  	// sem a quebra de linha no final
  	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite o nome do aluno(a): ")
	scanner.Scan()
  	nomeAluno = scanner.Text()

	fmt.Print("Nota 1: ")
	fmt.Scanln(&nota1)

	fmt.Print("Nota 2: ")
	fmt.Scanln(&nota2)

	fmt.Print("Nota 3: ")
	fmt.Scanln(&nota3)

	fmt.Printf("MÉDIA DE %s: %.2f", nomeAluno, calcMediaPonderada(nota1, nota2, nota3))
}