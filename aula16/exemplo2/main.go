package main

// Importa os pacotes necessários.
// fmt: utilizado para imprimir mensagens no terminal e escrever respostas HTTP.
// net/http: fornece os recursos para criar um servidor web.
import (
	"fmt"
	"net/http"
)

// Função principal da aplicação.
// É o ponto de entrada do programa.
func main() {

	// Registra a rota principal ("/") e associa
	// à função responsável por atendê-la.
	http.HandleFunc("/", home)

	// Registra a rota "/sobre".
	http.HandleFunc("/sobre", sobre)

	// Registra a rota "/contato".
	http.HandleFunc("/contato", contato)

	// Exibe uma mensagem informando que o servidor foi iniciado.
	fmt.Println("Servidor iniciado em http://localhost:8080")

	// Inicia o servidor HTTP na porta 8080.
	// Caso ocorra algum erro, ele será exibido no terminal.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}

// home é a função responsável por atender as
// requisições da página inicial.
//
// Parâmetros:
// w -> Permite enviar uma resposta ao navegador.
// r -> Contém todas as informações da requisição HTTP.
func home(w http.ResponseWriter, r *http.Request) {

	// Envia o texto "Página inicial" para o navegador.
	fmt.Fprintf(w, "Página inicial")
}

// sobre atende as requisições da página "/sobre".
func sobre(w http.ResponseWriter, r *http.Request) {

	// Envia o conteúdo da página Sobre.
	fmt.Fprintf(w, "Página sobre")
}

// contato atende as requisições da página "/contato".
func contato(w http.ResponseWriter, r *http.Request) {

	// Envia o conteúdo da página Contato.
	fmt.Fprintf(w, "Página de contato")
}
