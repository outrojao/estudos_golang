// Neste exercício para a prática da linguagem GoLang você deverá usar o laço for para desenhar o famoso padrão de diamante de estrelas. Você pode também usar o laço while, se assim você o desejar.
// O programa deverá pedir que o usuário informe a quantidade de linhas que marcará a metade do diamante.

package main

import "fmt"

func criaDiamante(qtdLinhas int) {
	qtdEspacos := qtdLinhas
	var qtdDiamantes int = 1

	for range qtdLinhas - 1 {
		for range qtdEspacos {
			fmt.Print(" ")
		}
		for range qtdDiamantes {
			fmt.Print("*")
		}
		qtdDiamantes += 2
		qtdEspacos--
		fmt.Println()
	}

	for range qtdLinhas {
		for range qtdEspacos {
			fmt.Print(" ")
		}
		for range qtdDiamantes {
			fmt.Print("*")
		}
		qtdDiamantes -= 2
		qtdEspacos++
		fmt.Println()
	}
}

func main() {
	fmt.Print("Informe a quantidade de linhas do diamante: ")
	var qtdLInhas int
	fmt.Scanln(&qtdLInhas)
	criaDiamante(qtdLInhas)
}

// CORREÇÃO:
/*
func criaDiamante(qtdLinhas int) {
    total := qtdLinhas*2 - 1 // altura total do diamante

    for i := 0; i < total; i++ {
        // calcula a largura da linha
        var estrelas int
        if i < qtdLinhas {
            estrelas = 2*i + 1
        } else {
            j := total - i - 1
            estrelas = 2*j + 1
        }

        espacos := (total - estrelas) / 2

        // imprime os espaços
        for s := 0; s < espacos; s++ {
            fmt.Print(" ")
        }
        // imprime as estrelas
        for e := 0; e < estrelas; e++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}
*/
// ERROS: Iterando int com range, Repetição de lógica, Cálculo de estrelas e espaços, estilo de programação procedural

// for i := 0; i < n; i++ é o jeito CORRETO em Go de repetir algo n vezes.

/*
total := qtdLinhas*2 - 1
for i := 0; i < total; i++ {
    ...
}
	DEIXA TUDO EM UM LOOP SÓ COM O TOTAL DE LINHAS DO DIAMANTE COMPLETO, NÃO DIVIDIDO EM DOIS COMO ANTERIORMENTE
*/

/*
var estrelas int
if i < qtdLinhas {
    estrelas = 2*i + 1
} else {
    j := total - i - 1
    estrelas = 2*j + 1
}

espacos := (total - estrelas) / 2
for s := 0; s < espacos; s++ {
    fmt.Print(" ")
}
	FAZ O CÁLCULO DE ESTRELAS E ESPAÇOS DIRETO NO LOOP SEM GERAR UM ESTADO MUTAVEL EM CADA ITERAÇÃO
*/
