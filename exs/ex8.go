// Escrever um programa em Golang para ler uma temperatura em graus Celsius e apresentá-la convertida em graus Fahrenheit. 
// A fórmula de conversão é: F = ((C * 9) / 5) + 32, sendo F a temperatura em Fahrenheit e C a temperatura em Celsius.

package main
import "fmt"

func convFahrenheit(celsius float64) (fahrenheit float64){
	fahrenheit = ((celsius * 9) / 5) + 32
	return
}

func main(){
	var celsius float64
	fmt.Print("Informe a temperatura em Celsius: ")
	fmt.Scanln(&celsius)
	fahrenheit := convFahrenheit(celsius)
	fmt.Println(fahrenheit)
}