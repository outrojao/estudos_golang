// Escreva um programa Go que leia o peso de uma pessoa na Terra e retorne o seu peso na Lua.
// peso_lua = (peso_terra / 9.81) * 1.622

package main
import "fmt"

func calcPesoLua(pesoTerra float64) (pesoLua float64){
	pesoLua = (pesoTerra / 9.81) * 1.622
	return
}

func main(){
	var pesoTerra float64
	fmt.Print("Digite seu peso na Terra: ")
	fmt.Scanln(&pesoTerra)
	fmt.Println(calcPesoLua(pesoTerra))
}