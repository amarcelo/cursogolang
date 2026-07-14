package main

// Importação das bibliotecas utilizadas pelo programa.
import (
	"context" // Permite criar contextos para controlar e cancelar tarefas.
	"fmt"     // Utilizado para imprimir mensagens na tela.
	"time"    // Utilizado para trabalhar com tempo (esperas, cronômetros, etc.).
)

// Função que representa uma tarefa que ficará executando até receber
// um comando de cancelamento.
func tarefa(ctx context.Context) {

	// Loop infinito.
	// A tarefa continuará executando até receber um sinal para parar.
	for {

		// O select aguarda algum evento acontecer.
		select {

		// Caso o contexto seja cancelado...
		case <-ctx.Done():

			// Exibe uma mensagem informando que a tarefa foi interrompida.
			fmt.Println("Tarefa cancelada!")

			// Encerra imediatamente a função.
			return

		// Caso não exista nenhum cancelamento...
		default:

			// Continua executando normalmente.
			fmt.Println("Executando...")

			// Aguarda 1 segundo antes de repetir o laço.
			// Isso serve apenas para não imprimir milhares de mensagens por segundo.
			time.Sleep(time.Second)
		}
	}
}

func main() {

	// Cria um contexto "pai" (Background)
	// juntamente com uma função chamada cancel().
	//
	// ctx -> será compartilhado entre as goroutines.
	// cancel -> poderá interromper todas as tarefas que usam esse contexto.
	ctx, cancel := context.WithCancel(context.Background())

	// Inicia uma nova goroutine executando a função tarefa().
	//
	// A palavra "go" faz a função rodar em paralelo ao restante do programa.
	go tarefa(ctx)

	// O programa principal espera 3 segundos.
	//
	// Enquanto isso, a goroutine continua imprimindo:
	// Executando...
	// Executando...
	// Executando...
	time.Sleep(3 * time.Second)

	// Após 3 segundos chamamos a função cancel().
	//
	// Isso envia um sinal para todo mundo que estiver usando o contexto.
	cancel()

	// Espera mais 1 segundo apenas para dar tempo da goroutine
	// perceber o cancelamento e imprimir:
	//
	// Tarefa cancelada!
	//
	// Sem essa espera, o programa principal poderia terminar antes da
	// goroutine conseguir mostrar a mensagem.
	time.Sleep(time.Second)
}
