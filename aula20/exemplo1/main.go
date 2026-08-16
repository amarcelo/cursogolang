package main

// ======================================================
// IMPORTAÇÃO DOS PACOTES
// ======================================================

import (
	"bufio"   // Permite ler dados digitados pelo usuário
	"errors"  // Permite trabalhar e comparar erros
	"fmt"     // Permite imprimir mensagens na tela
	"os"      // Permite trabalhar com arquivos e sistema operacional
	"strings" // Permite manipular textos
	"time"    // Permite trabalhar com data e hora
)

// ======================================================
// FUNÇÃO PRINCIPAL
// ======================================================
//
// Todo programa executável em Go começa pela função main().
//
// Neste programa, main() é responsável por:
// 1. Criar o leitor do teclado
// 2. Mostrar o menu
// 3. Receber a opção escolhida
// 4. Chamar a função correspondente
//
// ======================================================

func main() {

	// Criamos um Scanner.
	//
	// Scanner é uma ferramenta que permite ler dados.
	//
	// os.Stdin representa a entrada padrão do computador,
	// que normalmente é o TECLADO.
	//
	// Portanto:
	//
	// scanner = ferramenta para ler o que o usuário digitar.

	scanner := bufio.NewScanner(os.Stdin)

	// Criamos um loop infinito.
	//
	// "for" sem condição significa:
	//
	// REPITA PARA SEMPRE.
	//
	// O programa ficará mostrando o menu até que
	// o usuário escolha a opção 0.

	for {

		// Mostramos o menu na tela.

		fmt.Println("\n==============================")
		fmt.Println("     GERENCIADOR DE ARQUIVOS")
		fmt.Println("==============================")
		fmt.Println("1 - Criar arquivo")
		fmt.Println("2 - Adicionar dados ao arquivo")
		fmt.Println("3 - Apagar arquivo")
		fmt.Println("0 - Sair")
		fmt.Println("==============================")

		// Print é diferente de Println.
		//
		// Print NÃO pula para a próxima linha.
		//
		// Isso permite que o usuário digite sua opção
		// na mesma linha.

		fmt.Print("Escolha uma opção: ")

		// Scan() espera o usuário digitar alguma coisa
		// e pressionar ENTER.

		scanner.Scan()

		// Text() pega aquilo que foi digitado.
		//
		// A informação recebida será uma string.

		opcao := scanner.Text()

		// switch verifica o valor da variável opcao.
		//
		// Dependendo do valor, executaremos uma função.

		switch opcao {

		// Se o usuário digitou 1:
		case "1":

			// Chamamos a função responsável
			// pela criação do arquivo.

			criarArquivo(scanner)

		// Se o usuário digitou 2:
		case "2":

			// Chamamos a função responsável
			// por adicionar dados.

			adicionarDados(scanner)

		// Se o usuário digitou 3:
		case "3":

			// Chamamos a função responsável
			// por apagar arquivos.

			apagarArquivo(scanner)

		// Se o usuário digitou 0:
		case "0":

			fmt.Println("Programa encerrado.")

			// Registramos no log que o programa
			// foi encerrado.

			gravarLog("Programa encerrado")

			// return encerra a função main().
			//
			// Como main() é a função principal,
			// o programa termina.

			return

		// Qualquer outra coisa digitada
		// cairá no default.

		default:

			fmt.Println("Opção inválida!")
		}
	}
}

// ======================================================
// FUNÇÃO criarArquivo
// ======================================================
//
// Esta função recebe o Scanner como parâmetro.
//
// O Scanner é utilizado para perguntar ao usuário
// qual será o nome do arquivo.
//
// Exemplo:
//
// dados.txt
//
// A função verifica se o arquivo já existe antes
// de tentar criá-lo.
//
// ======================================================

func criarArquivo(scanner *bufio.Scanner) {

	// Perguntamos o nome do arquivo.

	fmt.Print("Digite o nome do arquivo: ")

	// Esperamos o usuário digitar.

	scanner.Scan()

	// Pegamos o texto digitado.
	//
	// strings.TrimSpace() remove espaços desnecessários
	// no início e no final.
	//
	// Por exemplo:
	//
	// "   dados.txt   "
	//
	// torna-se:
	//
	// "dados.txt"

	nome := strings.TrimSpace(scanner.Text())

	// Verificamos se o usuário não digitou nada.

	if nome == "" {

		fmt.Println("Nome inválido!")

		// Registramos a tentativa no log.

		gravarLog("Tentativa de criar arquivo sem nome")

		return
	}

	// --------------------------------------------------
	// VERIFICANDO SE O ARQUIVO EXISTE
	// --------------------------------------------------

	// os.Stat() tenta obter informações sobre
	// um arquivo ou diretório.
	//
	// Ele retorna duas informações:
	//
	// FileInfo
	// erro
	//
	// Neste exemplo não precisamos do FileInfo.
	//
	// Por isso utilizamos "_".
	//
	// O "_" significa:
	//
	// "Ignore este valor."

	_, err := os.Stat(nome)

	// Se err for nil, significa que NÃO aconteceu erro.
	//
	// Portanto, o arquivo foi encontrado.

	if err == nil {

		fmt.Println("O arquivo já existe!")

		gravarLog(
			"Tentativa de criar arquivo já existente: " + nome,
		)

		return
	}

	// --------------------------------------------------
	// USANDO errors.Is()
	// --------------------------------------------------

	// Aqui sabemos que aconteceu algum erro.
	//
	// Mas queremos descobrir QUAL erro aconteceu.
	//
	// os.ErrNotExist representa:
	//
	// "arquivo ou diretório não existe".
	//
	// errors.Is() pergunta:
	//
	// "O erro recebido significa que o arquivo
	// não existe?"
	//
	// Temos um ! antes de errors.Is().
	//
	// ! significa NÃO.
	//
	// Portanto:
	//
	// if !errors.Is(...)
	//
	// significa:
	//
	// "SE NÃO for erro de arquivo inexistente..."

	if !errors.Is(err, os.ErrNotExist) {

		fmt.Println("Erro ao verificar arquivo:", err)

		gravarLog(
			"Erro ao verificar arquivo: " + nome,
		)

		return
	}

	// Se chegamos até aqui significa:
	//
	// 1. O arquivo não existe
	// 2. Podemos tentar criá-lo
	//
	// os.Create() cria um novo arquivo.

	arquivo, err := os.Create(nome)

	// Verificamos se aconteceu algum erro
	// durante a criação.

	if err != nil {

		fmt.Println("Erro ao criar arquivo:", err)

		gravarLog(
			"Erro ao criar arquivo: " + nome,
		)

		return
	}

	// O arquivo foi aberto pelo sistema operacional.
	//
	// Precisamos fechá-lo depois de usar.
	//
	// Como não vamos escrever nada nele neste momento,
	// podemos fechá-lo imediatamente.

	arquivo.Close()

	fmt.Println("Arquivo criado com sucesso!")

	// Registramos a operação.

	gravarLog(
		"Arquivo criado: " + nome,
	)
}

// ======================================================
// FUNÇÃO adicionarDados
// ======================================================
//
// Esta função adiciona uma mensagem ao FINAL
// de um arquivo existente.
//
// Exemplo:
//
// Arquivo:
//
// Primeira mensagem
//
// Depois adicionamos:
//
// Segunda mensagem
//
// Resultado:
//
// Primeira mensagem
// Segunda mensagem
//
// O conteúdo anterior NÃO será apagado.
//
// ======================================================

func adicionarDados(scanner *bufio.Scanner) {

	// Perguntamos qual arquivo será utilizado.

	fmt.Print("Digite o nome do arquivo: ")

	scanner.Scan()

	nome := strings.TrimSpace(scanner.Text())

	// --------------------------------------------------
	// VERIFICAMOS SE O ARQUIVO EXISTE
	// --------------------------------------------------

	_, err := os.Stat(nome)

	// Perguntamos:
	//
	// "O erro aconteceu porque o arquivo não existe?"

	if errors.Is(err, os.ErrNotExist) {

		fmt.Println("O arquivo não existe!")

		gravarLog(
			"Tentativa de escrever em arquivo inexistente: " + nome,
		)

		return
	}

	// Pode ter acontecido outro tipo de erro.
	//
	// Por exemplo:
	//
	// falta de permissão.

	if err != nil {

		fmt.Println("Erro ao acessar arquivo:", err)

		gravarLog(
			"Erro ao acessar arquivo: " + nome,
		)

		return
	}

	// --------------------------------------------------
	// PEDIMOS A MENSAGEM
	// --------------------------------------------------

	fmt.Print("Digite a mensagem que deseja adicionar: ")

	scanner.Scan()

	mensagem := scanner.Text()

	// --------------------------------------------------
	// ABRIMOS O ARQUIVO
	// --------------------------------------------------
	//
	// os.OpenFile() permite controlar COMO
	// queremos abrir o arquivo.
	//
	// Estamos utilizando:
	//
	// os.O_APPEND
	//
	// Adiciona novos dados ao FINAL do arquivo.
	//
	//
	// os.O_WRONLY
	//
	// Abre o arquivo somente para ESCRITA.
	//
	//
	// 0644
	//
	// Representa as permissões do arquivo
	// em sistemas Unix/Linux.

	arquivo, err := os.OpenFile(
		nome,
		os.O_APPEND|os.O_WRONLY,
		0644,
	)

	// Verificamos se conseguimos abrir o arquivo.

	if err != nil {

		fmt.Println("Erro ao abrir arquivo:", err)

		gravarLog(
			"Erro ao abrir arquivo para escrita: " + nome,
		)

		return
	}

	// --------------------------------------------------
	// DEFER
	// --------------------------------------------------
	//
	// defer significa:
	//
	// "Execute esta instrução quando a função terminar."
	//
	// Portanto:
	//
	// defer arquivo.Close()
	//
	// significa:
	//
	// "Quando adicionarDados() terminar,
	// feche o arquivo."

	defer arquivo.Close()

	// --------------------------------------------------
	// ESCREVENDO NO ARQUIVO
	// --------------------------------------------------
	//
	// WriteString() escreve uma string no arquivo.
	//
	// "\n" representa uma quebra de linha.

	_, err = arquivo.WriteString(
		mensagem + "\n",
	)

	// Verificamos se aconteceu algum erro
	// durante a escrita.

	if err != nil {

		fmt.Println(
			"Erro ao escrever no arquivo:",
			err,
		)

		gravarLog(
			"Erro ao escrever no arquivo: " + nome,
		)

		return
	}

	fmt.Println(
		"Mensagem adicionada com sucesso!",
	)

	// Registramos a operação.

	gravarLog(
		"Mensagem adicionada ao arquivo: " + nome,
	)
}

// ======================================================
// FUNÇÃO apagarArquivo
// ======================================================
//
// Esta função:
//
// 1. Pergunta qual arquivo deve ser apagado
// 2. Verifica se ele existe
// 3. Pede confirmação
// 4. Apaga o arquivo
// 5. Registra a operação no log
//
// ======================================================

func apagarArquivo(scanner *bufio.Scanner) {

	// Perguntamos o nome do arquivo.

	fmt.Print(
		"Digite o nome do arquivo que deseja apagar: ",
	)

	scanner.Scan()

	nome := strings.TrimSpace(scanner.Text())

	// --------------------------------------------------
	// VERIFICAMOS SE O ARQUIVO EXISTE
	// --------------------------------------------------

	_, err := os.Stat(nome)

	// errors.Is() verifica se o erro significa
	// que o arquivo não existe.

	if errors.Is(err, os.ErrNotExist) {

		fmt.Println("O arquivo não existe!")

		gravarLog(
			"Tentativa de apagar arquivo inexistente: " + nome,
		)

		return
	}

	// Se aconteceu outro erro...

	if err != nil {

		fmt.Println(
			"Erro ao verificar arquivo:",
			err,
		)

		gravarLog(
			"Erro ao verificar arquivo: " + nome,
		)

		return
	}

	// --------------------------------------------------
	// PEDIMOS CONFIRMAÇÃO
	// --------------------------------------------------
	//
	// Neste ponto sabemos que o arquivo existe.
	//
	// Antes de apagar, perguntamos ao usuário.

	fmt.Printf(
		"Tem certeza que deseja apagar '%s'? (s/n): ",
		nome,
	)

	scanner.Scan()

	// Pegamos a resposta.
	//
	// TrimSpace remove espaços.
	//
	// ToLower transforma tudo em minúsculo.
	//
	// Portanto:
	//
	// S
	// s
	//
	// serão tratados da mesma maneira.

	resposta := strings.ToLower(
		strings.TrimSpace(scanner.Text()),
	)

	// Se a resposta NÃO for "s",
	// cancelamos a exclusão.

	if resposta != "s" {

		fmt.Println("Operação cancelada.")

		gravarLog(
			"Exclusão cancelada pelo usuário: " + nome,
		)

		return
	}

	// --------------------------------------------------
	// APAGANDO O ARQUIVO
	// --------------------------------------------------

	// os.Remove() solicita ao sistema operacional
	// que remova o arquivo.

	err = os.Remove(nome)

	// Verificamos se aconteceu algum erro.

	if err != nil {

		fmt.Println(
			"Erro ao apagar arquivo:",
			err,
		)

		gravarLog(
			"Erro ao apagar arquivo: " + nome,
		)

		return
	}

	fmt.Println(
		"Arquivo apagado com sucesso!",
	)

	// Registramos a operação.

	gravarLog(
		"Arquivo apagado: " + nome,
	)
}

// ======================================================
// FUNÇÃO gravarLog
// ======================================================
//
// Esta função é responsável por registrar
// tudo o que acontece no programa.
//
// Ela recebe uma mensagem:
//
// gravarLog("Arquivo criado: teste.txt")
//
// E adiciona essa mensagem ao arquivo:
//
// log.txt
//
// Também adicionamos a data e a hora.
//
// ======================================================

func gravarLog(mensagem string) {

	// --------------------------------------------------
	// ABRINDO O LOG
	// --------------------------------------------------
	//
	// Utilizamos os.OpenFile().
	//
	// Existem três opções importantes:
	//
	// os.O_CREATE
	//
	// Se log.txt NÃO existir:
	//
	//     CRIE o arquivo.
	//
	//
	// os.O_APPEND
	//
	// Se o arquivo já tiver informações:
	//
	//     NÃO apague.
	//
	// Coloque a nova informação no FINAL.
	//
	//
	// os.O_WRONLY
	//
	// Abra o arquivo somente para escrita.

	arquivo, err := os.OpenFile(
		"log.txt",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)

	// Verificamos se conseguimos abrir
	// ou criar o arquivo log.txt.

	if err != nil {

		fmt.Println(
			"Erro ao abrir log:",
			err,
		)

		return
	}

	// Garantimos que o arquivo será fechado
	// quando a função terminar.

	defer arquivo.Close()

	// --------------------------------------------------
	// DATA E HORA
	// --------------------------------------------------

	// time.Now() pega a data e hora atual
	// do computador.

	agora := time.Now()

	// --------------------------------------------------
	// MONTANDO A LINHA DO LOG
	// --------------------------------------------------
	//
	// fmt.Sprintf() monta uma string formatada.
	//
	// Exemplo do resultado:
	//
	// [14/08/2026 16:45:32] Arquivo criado: dados.txt
	//
	// O primeiro %s recebe a data.
	// O segundo %s recebe a mensagem.

	linha := fmt.Sprintf(
		"[%s] %s\n",
		agora.Format("02/01/2006 15:04:05"),
		mensagem,
	)

	// --------------------------------------------------
	// GRAVANDO NO LOG
	// --------------------------------------------------

	// Escrevemos a linha dentro do log.txt.

	_, err = arquivo.WriteString(linha)

	// Verificamos se aconteceu algum erro.

	if err != nil {

		fmt.Println(
			"Erro ao gravar log:",
			err,
		)
	}
}
