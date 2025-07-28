// Escreva um programa GoLang que leia dois números e mostre a sua soma. Os dois números deverão ser informados pelo usuário e você deverá efetuar a leitura como dois inteiros e mostrar a soma também como um inteiro. Os valores informados pelo usuário podem ser positivos ou negativos.

package main
import "fmt"

func soma(x, y int) int {
	return x + y
}

func main() {
	var num1, num2 int
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	fmt.Println("RESULTADO:", soma(num1, num2))
}