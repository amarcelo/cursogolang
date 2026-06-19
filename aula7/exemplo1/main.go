package main

import "fmt"

func main() {

	// Cria um map onde:
	// string = nome da pessoa
	// string = telefone da pessoa
	contatos := make(map[string]string)

	// Adicionando contatos
	contatos["Maria"] = "9999-1111"
	contatos["João"] = "9999-2222"
	contatos["Pedro"] = "9999-3333"
	contatos["Ana"] = "9999-5555"

	// Consultando um contato
	fmt.Println("Telefone da Maria:", contatos["Maria"])

	// Alterando um contato existente
	contatos["Pedro"] = "9888-3333"
	contatos["Ana"] = "7777-2222"

	fmt.Println("Novo telefone do Pedro:", contatos["Pedro"])
	fmt.Println("Novo telefone da Aba:", contatos["Ana"])
	fmt.Println("Novo telefone do Altair:", contatos["Altair"])

	// Percorrendo todos os contatos
	fmt.Println("\nLista de contatos:")

	for nome, telefone := range contatos {
		fmt.Printf("%s -> %s\n", nome, telefone)
	}
}
