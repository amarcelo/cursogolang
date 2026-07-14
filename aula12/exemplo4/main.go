package main

// Importação das bibliotecas usadas pelo programa.
import (
	"context" // Permite criar um contexto para controlar e cancelar a goroutine.
	"fmt"     // Permite imprimir mensagens no terminal.
	"time"    // Permite trabalhar com tempo e pausas.
)

/*
A função worker representa uma tarefa que será executada
em segundo plano.

Ela recebe um parâmetro chamado ctx, do tipo context.Context.

Esse contexto será usado para avisar ao worker quando ele
deve encerrar seu trabalho.
*/
func worker(ctx context.Context) {

	/*
		O laço for sem condição cria uma repetição infinita.

		Isso significa que o worker continuará trabalhando
		até receber algum comando para parar.
	*/
	for {

		/*
			O select é usado para observar operações envolvendo canais.

			Neste programa, ele verifica duas possibilidades:

			1. O contexto foi cancelado.
			2. O contexto ainda está ativo e o worker deve continuar.
		*/
		select {

		/*
			ctx.Done() retorna um canal.

			Enquanto o contexto estiver ativo, esse canal permanece aberto
			e não libera esta opção do select.

			Quando a função cancel() for chamada, o canal será fechado.

			O trecho:

			<-ctx.Done()

			significa que o worker está verificando se recebeu
			o sinal de cancelamento.
		*/
		case <-ctx.Done():

			// Esta mensagem é exibida quando o contexto é cancelado.
			fmt.Println("Worker encerrado")

			/*
				O return encerra imediatamente a função worker.

				Como essa função está sendo executada em uma goroutine,
				o return também encerra essa goroutine.
			*/
			return

		/*
			O bloco default é executado quando nenhuma das outras
			opções do select está pronta.

			Neste caso, ele será executado enquanto o contexto
			ainda não tiver sido cancelado.
		*/
		default:

			// Mostra que o worker ainda está realizando seu trabalho.
			fmt.Println("Processando")

			/*
				Faz a goroutine esperar durante 1 segundo.

				Essa pausa evita que o programa imprima milhares
				de mensagens por segundo e simula uma tarefa demorada.
			*/
			time.Sleep(time.Second)
		}
	}
}

func main() {

	/*
		context.Background() cria um contexto inicial.

		Esse contexto não possui cancelamento, prazo ou valores.
		Ele serve como contexto pai para o novo contexto que será criado.
	*/

	/*
		context.WithCancel cria um contexto que pode ser cancelado
		manualmente.

		A função retorna dois valores:

		ctx:
		O contexto que será entregue ao worker.

		cancel:
		Uma função usada para enviar o sinal de cancelamento.
	*/
	ctx, cancel := context.WithCancel(context.Background())

	/*
		Inicia a função worker em uma nova goroutine.

		A palavra go faz com que worker(ctx) execute de forma concorrente,
		enquanto a função main continua executando suas próximas linhas.
	*/
	go worker(ctx)

	/*
		A função main espera durante 5 segundos.

		Enquanto isso, a goroutine worker continua executando e imprimindo:

		Processando
		Processando
		Processando
		...
	*/
	time.Sleep(6 * time.Second)

	/*
		Chama a função de cancelamento.

		Essa função fecha o canal retornado por ctx.Done().

		Quando o worker verificar novamente o select, ele entrará no caso:

		case <-ctx.Done()

		e encerrará sua execução.
	*/
	cancel()

	/*
		A função main espera mais 1 segundo.

		Essa espera serve para dar tempo de a goroutine perceber
		o cancelamento, imprimir "Worker encerrado" e finalizar.

		Sem essa pausa, a função main poderia terminar imediatamente,
		encerrando todo o programa antes de a goroutine imprimir
		a mensagem final.
	*/
	time.Sleep(time.Second)
}
