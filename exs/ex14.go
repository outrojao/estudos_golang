/*
Um losango é um quadrilátero equilátero, ou seja, é um polígono formado por quatro lados de igual comprimento. Um losango é também um paralelogramo. Alguns autores exigem ainda que nenhum dos ângulos do quadrilátero seja reto para que ele seja considerado um losango.

A área (em metros quadrados) de um losango pode ser calculada usando-se a seguinte fórmula: Al = D1 * D2 / 2
*/

package main

import "fmt"

func calcAreaLosango(diagonalMaior, diagonalMenor int) (area int) {
	area = (diagonalMaior * diagonalMenor) / 2
	return
}

func main() {
	fmt.Print("Diagonal maior: ")
	var diagonalMaior int
	fmt.Scanln(&diagonalMaior)

	fmt.Print("Diagonal menor: ")
	var diagonalMenor int
	fmt.Scanln(&diagonalMenor)

	fmt.Printf("ÁREA DO LOSANGO: %d metros quadrados", calcAreaLosango(diagonalMaior, diagonalMenor))
}
