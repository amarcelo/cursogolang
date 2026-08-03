package main

// Importa os pacotes necessários:
// fmt: utilizado para imprimir mensagens no terminal e escrever respostas HTTP.
// net/http: fornece funcionalidades para criação de servidores web.
import (
	"fmt"
	"net/http"
)

// Função principal do programa.
// É o ponto de entrada da aplicação.
func main() {

	// Registra uma rota ("/") e associa essa rota
	// à função homeHandler.
	// Sempre que um usuário acessar a URL raiz,
	// essa função será executada.
	http.HandleFunc("/", homeHandler)

	// Exibe uma mensagem no terminal informando
	// que o servidor está em execução.
	fmt.Println("Servidor rodando em http://localhost:8080")

	// Inicia o servidor HTTP na porta 8080.
	// O segundo parâmetro (nil) indica que será utilizado
	// o roteador padrão (DefaultServeMux).
	http.ListenAndServe(":8080", nil)
}

// Função responsável por tratar as requisições
// recebidas na rota "/".
//
// Parâmetros:
// w -> ResponseWriter: utilizado para enviar a resposta ao navegador.
// r -> Request: contém todas as informações da requisição recebida.
func homeHandler(w http.ResponseWriter, r *http.Request) {

	// Escreve a mensagem "Olá Mundo!" na resposta HTTP,
	// que será exibida no navegador do usuário.
	fmt.Fprintf(w, "Bem vido ao meu servidor web em GO!")
}
