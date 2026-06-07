package main

import "fmt"

func main() {

	var nota float64 = 7

	switch {

	case nota >= 9:
		fmt.Println("Excelente")

	case nota >= 7:
		fmt.Println("Aprovado")

	case nota >= 5:
		fmt.Println("Recuperação")

	default:
		fmt.Println("Reprovado")
	}
}
