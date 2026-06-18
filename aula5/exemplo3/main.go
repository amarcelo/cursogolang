package main

import "fmt"

func criarContador() func() int {

	contador := 0

	return func() int {
		contador++
		return contador
	}
}

func main() {

	contar := criarContador()

	fmt.Println(contar())
	fmt.Println(contar())
	fmt.Println(contar())
	fmt.Println(contar())

}
