package main

import (
	"context"
	"fmt"
	"time"
)

/*
A função baixarArquivo simula o download de um arquivo.

Ela recebe dois parâmetros:

 1. ctx context.Context
    É o contexto usado para avisar quando o download deve ser cancelado.

 2. nomeArquivo string
    É o nome do arquivo que será baixado.

A função será executada dentro de uma goroutine.
*/
func baixarArquivo(ctx context.Context, nomeArquivo string) {

	/*
		A variável progresso começa com zero.

		Ela representa a porcentagem já baixada do arquivo.
	*/
	progresso := 0

	/*
		Este laço continuará executando enquanto o progresso
		for menor ou igual a 100%.
	*/
	for progresso <= 100 {

		/*
			O select verifica se aconteceu algum evento importante.

			Neste caso, ele observa duas possibilidades:

			1. O contexto foi cancelado.
			2. O contexto ainda está ativo e o download pode continuar.
		*/
		select {

		/*
			ctx.Done() retorna um canal que é fechado
			quando o contexto é cancelado.

			Quando a função cancel() for chamada na main,
			este case será executado.
		*/
		case <-ctx.Done():

			// Informa que o download foi interrompido.
			fmt.Println("\nDownload cancelado!")

			// Exibe o motivo do encerramento do contexto.
			fmt.Println("Motivo:", ctx.Err())

			/*
				O return encerra a função baixarArquivo.

				Como essa função está rodando em uma goroutine,
				o return também encerra essa goroutine.
			*/
			return

		/*
			O bloco default é executado quando o contexto
			ainda não foi cancelado.
		*/
		default:

			// Exibe o nome do arquivo e o progresso atual.
			fmt.Printf(
				"Baixando %s: %d%%\n",
				nomeArquivo,
				progresso,
			)

			/*
				Simula o tempo necessário para baixar
				uma parte do arquivo.

				A goroutine espera 1 segundo antes de continuar.
			*/
			time.Sleep(time.Second)

			/*
				Aumenta o progresso em 10%.

				Na próxima repetição, será exibido um valor maior.
			*/
			progresso += 10
		}
	}

	/*
		Esta mensagem só será exibida se o progresso chegar a 100%
		sem que o contexto seja cancelado.
	*/
	fmt.Println("\nDownload concluído com sucesso!")
}

func main() {

	/*
		context.Background() cria um contexto inicial.

		context.WithCancel cria um novo contexto
		que pode ser cancelado manualmente.

		A função retorna:

		ctx:
		O contexto que será enviado para a função de download.

		cancel:
		A função que será usada para cancelar o download.
	*/
	ctx, cancel := context.WithCancel(context.Background())

	/*
		defer cancel() garante que o contexto será encerrado
		quando a função main terminar.

		Isso ajuda a liberar recursos internos.
	*/
	defer cancel()

	/*
		Inicia o download em uma nova goroutine.

		A palavra go faz com que baixarArquivo execute
		ao mesmo tempo que a função main.
	*/
	go baixarArquivo(ctx, "curso-go.zip")

	/*
		A função main espera 5 segundos.

		Durante esse tempo, o download continua acontecendo
		na outra goroutine.
	*/
	time.Sleep(6 * time.Second)

	/*
		Depois de 5 segundos, simulamos que o usuário
		clicou no botão "Cancelar download".
	*/
	fmt.Println("\nUsuário clicou em cancelar.")

	/*
		Chama a função de cancelamento.

		Isso fecha o canal retornado por ctx.Done().

		Na próxima verificação do select, a função baixarArquivo
		perceberá o cancelamento e será encerrada.
	*/
	cancel()

	/*
		A função main espera mais 1 segundo.

		Essa pausa dá tempo para a goroutine perceber
		o cancelamento e imprimir as mensagens finais.
	*/
	time.Sleep(time.Second)

	// Informa que o programa principal foi encerrado.
	fmt.Println("Programa finalizado.")
}
