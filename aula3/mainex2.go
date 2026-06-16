package main

import (
	"fmt"
	"math"
)

var nome string = "Maria"
var soma int = 10

//Executando o init

func init() {
	fmt.Println("Iniciando programa...")
}

func main() {
	nome := "Antonio"
	raiz := math.Sqrt(16)
	soma := raiz + 5

	fmt.Println("Nome Local:", nome)
	fmt.Println("Raiz quadrada de 16:", raiz)
	fmt.Println("Soma da raiz quadrada local com 5:", soma)

	imprimeglobal()
}

func imprimeglobal() {
	fmt.Println("Nome Global: ", nome)
	fmt.Println("Soma Global: ", soma)
}
