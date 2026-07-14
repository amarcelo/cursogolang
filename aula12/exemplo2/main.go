package main

// Importamos as bibliotecas que serão utilizadas no programa.
import (
	"context" // Biblioteca usada para controlar tempo, cancelamento e encerramento de tarefas.
	"fmt"     // Biblioteca usada para imprimir mensagens no terminal.
	"time"    // Biblioteca usada para trabalhar com duração e passagem do tempo.
)

func main() {

	/*
		context.WithTimeout cria um contexto que será cancelado
		automaticamente depois de um período determinado.

		Neste exemplo, o tempo definido é de 3 segundos.

		A função context.WithTimeout retorna dois valores:

		1. ctx
		   É o contexto criado. Ele carrega a informação de que existe
		   um limite de tempo de 3 segundos.

		2. cancel
		   É uma função que permite cancelar manualmente o contexto,
		   mesmo antes de os 3 segundos terminarem.
	*/
	ctx, cancel := context.WithTimeout(
		context.Background(), // Contexto inicial que servirá como contexto pai.
		3*time.Second,        // Define que o contexto deve durar no máximo 3 segundos.
	)

	/*
		defer faz com que a função cancel() seja executada
		quando a função main estiver terminando.

		É uma boa prática chamar cancel(), mesmo quando o contexto
		já possui um tempo limite automático.

		Isso ajuda a liberar recursos internos utilizados pelo contexto,
		como temporizadores e estruturas de controle.
	*/
	defer cancel()

	/*
		ctx.Done() retorna um canal.

		Esse canal permanece bloqueado enquanto o contexto estiver ativo.

		Quando os 3 segundos terminarem, o contexto será cancelado
		automaticamente e o canal será fechado.

		O operador <- espera receber um sinal desse canal.

		Portanto, esta linha faz o programa parar aqui e esperar
		até que o tempo limite do contexto seja atingido.
	*/
	<-ctx.Done()

	/*
		Esta linha somente será executada depois que o contexto terminar.

		Como o contexto foi criado com um limite de 3 segundos,
		a mensagem será exibida aproximadamente 3 segundos
		depois do início do programa.
	*/
	fmt.Println("Tempo esgotado")
}
