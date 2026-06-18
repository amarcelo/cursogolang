package main

import "fmt"

func dividir(a, b float64) (float64, string) {

	if b == 0 {
		return 0, "Erro: divisão por zero"
	}

	return a / b, "OK"
}

func main() {

	resultado, status := dividir(10, 0)

	fmt.Println("Resultado:", resultado)
	fmt.Println("Status:", status)

}
