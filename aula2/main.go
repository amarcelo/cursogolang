package main

import (
	"fmt"
	"math"
)

//Executando o init

func init() {
	fmt.Println("Iniciando programa...")
}

func main() {
	var nome string = "Antonio"
	var raiz float64 = math.Sqrt(16)
	var soma float64 = raiz + 5

	fmt.Println("Nome:", nome)
	fmt.Println("Raiz quadrada de 16:", raiz)
	fmt.Println("Soma da raiz quadrada com 5:", soma)
}
