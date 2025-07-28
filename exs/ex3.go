// Escreva um programa GoLang que receberá a duração de um evento expresso em segundos e exiba-o expresso em horas, minutos e segundos.

package main
import "fmt"

func calcTemp(segundos int) string {
	horas := segundos / 3600
    minutos := (segundos % 3600) / 60
    segundos = segundos % 60

    return fmt.Sprintf("%02d:%02d:%02d", horas, minutos, segundos)
}

func main() {
	var segundos int
	fmt.Print("Duração em segundos: ")
	fmt.Scanln(&segundos)
	fmt.Println(calcTemp(segundos))
}