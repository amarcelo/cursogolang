package main

import "fmt"

func main() {

	var senha string
	var tentativas int = 0

INICIO:

	fmt.Print("Digite a senha: ")
	fmt.Scanln(&senha)

	if senha == "golang123" {
		fmt.Println("Acesso liberado!")
		goto FIM
	}

	tentativas++

	fmt.Println("Senha incorreta.")

	if tentativas < 3 {
		fmt.Printf("Tentativa %d de 3\n", tentativas)
		goto INICIO
	}

	fmt.Println("Número máximo de tentativas atingido.")

FIM:
	fmt.Println("Programa encerrado.")
}
