// Escreva um programa GoLang que solicita ao usuário dois números inteiros e mostre a multiplicação dos dois valores, ou seja, o primeiro valor multiplicado pelo segundo.

package main
import "fmt"

func multi(x, y int) int {
	return x * y
}

func main() {
	var num1, num2 int
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	fmt.Println("RESULTADO:", multi(num1, num2))
}