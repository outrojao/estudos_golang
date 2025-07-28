package main
import "fmt"

func encontraParEImpar(numLista []int) (resultadoLista []string) {
    for posi := range numLista {
        if numLista[posi]%2 == 0 {
            resultadoLista = append(resultadoLista, "PAR")
        } else {
            resultadoLista = append(resultadoLista, "IMPAR")
        }
    }
    return
}

func main(){
	nums := []int{1,2,3,4,5}
	fmt.Println(encontraParEImpar(nums))
}