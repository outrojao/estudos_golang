// A empresa GÁS TOTAL paga uma gratificação de R$ 237,50 a todos os seus funcionários. Escreva um programa Go que pede o salário do funcionário e gera o novo salário com a gratificação.
// As variáveis que representam a gratificação, o salário atual e o novo salário deverão ser do tipo float ou double.

package main
import "fmt"

func gratificar(salario, gratificacao float64) (novoSalario float64) {
	novoSalario = salario + gratificacao
	return
}

func main() {
	var salario float64
	fmt.Print("Digite o valor do salário: ")
	fmt.Scanln(&salario)
	fmt.Println(gratificar(salario, 237.50))
}