// Faça um programa em Golang que receba dois números e no final mostre a soma, subtração, multiplicação e a divisão dos números lidos. Os números deverão ser informados pelo usuário.

package main
import "fmt"

func soma(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func multi(x, y int) int {
	return x * y
}

func divi(x, y int) int {
	return x / y
}


func main() {
	var num1, num2 int
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	fmt.Println("SOMA:", soma(num1, num2))
	fmt.Println("SUBTRAÇÃO:", sub(num1, num2))
	fmt.Println("MULTIPLICAÇÃO:", multi(num1, num2))
	fmt.Println("DIVISÃO:", divi(num1, num2))
}