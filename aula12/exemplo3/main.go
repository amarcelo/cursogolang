package main

// Importamos as bibliotecas que serão usadas no programa.
import (
	"context" // Fornece recursos para criar e transportar contextos.
	"fmt"     // Permite imprimir informações no terminal.
)

func main() {

	/*
		context.WithValue cria um novo contexto contendo
		uma informação no formato chave e valor.

		Neste exemplo:

		Chave: "usuario"
		Valor: "Antonio"

		É parecido com guardar uma informação em uma etiqueta:

		"usuario" -> "Antonio"

		Depois, qualquer função que receber esse contexto poderá
		tentar consultar essa informação.
	*/
	ctx := context.WithValue(

		/*
			context.Background() cria o contexto inicial.

			Ele funciona como o contexto pai do novo contexto.

			O contexto Background:

			- não possui valores;
			- não possui cancelamento;
			- não possui tempo limite;
			- geralmente é usado como ponto inicial de um programa.
		*/
		context.Background(),

		/*
			Este é o nome da chave usada para guardar a informação.

			A chave permite localizar o valor posteriormente.

			Neste caso, a chave é:

			"usuario"
		*/
		"usuario",

		/*
			Este é o valor associado à chave "usuario".

			O valor armazenado é:

			"Antonio"
		*/
		"Joao",
	)

	/*
		ctx.Value procura dentro do contexto um valor associado
		à chave informada.

		Neste caso, estamos procurando a chave:

		"usuario"

		Como essa chave foi armazenada anteriormente, o contexto
		retornará o valor:

		"Antonio"
	*/
	nome := ctx.Value("usuario")

	/*
		fmt.Println imprime o conteúdo da variável nome no terminal.

		Como nome recebeu o valor associado à chave "usuario",
		a saída será:

		Antonio
	*/
	fmt.Println(nome)
}
