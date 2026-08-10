package main

// Importamos os pacotes que vamos utilizar no programa.
import (
	"encoding/json" // Trabalha com JSON.
	"fmt"           // Permite mostrar mensagens no terminal.
	"net/http"      // Permite criar servidor e trabalhar com HTTP.
	"strconv"       // Permite converter texto para número.
)

// ----------------------------------------------------
// STRUCT PRODUTO
// ----------------------------------------------------

// Produto representa os dados de um produto.
//
// Uma struct serve para agrupar informações relacionadas.
//
// Neste caso, cada produto terá:
// ID
// Nome
// Preço
type Produto struct {

	// ID é o número identificador do produto.
	ID int `json:"id"`

	// Nome armazena o nome do produto.
	Nome string `json:"nome"`

	// Preco armazena o preço.
	//
	// Usamos float64 porque o preço pode ter casas decimais.
	Preco float64 `json:"preco"`
}

// ----------------------------------------------------
// LISTA DE PRODUTOS
// ----------------------------------------------------

// Aqui criamos uma lista de produtos.
//
// []Produto significa:
//
// "uma lista de elementos do tipo Produto"
//
// Estamos usando essa lista para simular um banco de dados.
//
// Quando o programa for fechado, os dados cadastrados
// durante a execução serão perdidos.
var produtos = []Produto{

	{
		ID:    1,
		Nome:  "Notebook",
		Preco: 4500,
	},

	{
		ID:    2,
		Nome:  "Mouse",
		Preco: 120,
	},

	{
		ID:    3,
		Nome:  "Teclado Mecânico",
		Preco: 350,
	},
}

// ----------------------------------------------------
// FUNÇÃO PRINCIPAL
// ----------------------------------------------------

func main() {

	// Criamos a rota /produtos.
	//
	// Essa rota será tratada pela função produtosHandler.
	//
	// Ela poderá aceitar:
	//
	// GET  /produtos
	// POST /produtos
	http.HandleFunc("/produtos", produtosHandler)

	// Criamos a rota /produto.
	//
	// Ela será utilizada para procurar um produto pelo ID.
	//
	// Exemplo:
	//
	// GET /produto?id=2
	http.HandleFunc("/produto", produtoPorIDHandler)

	// Mostra uma mensagem no terminal
	// para sabermos que o servidor iniciou.
	fmt.Println("Servidor em http://localhost:8080")

	// Inicia o servidor HTTP na porta 8080.
	//
	// O nil significa que estamos utilizando
	// o roteador padrão do pacote net/http.
	http.ListenAndServe(":8080", nil)
}

// ----------------------------------------------------
// HANDLER DA ROTA /produtos
// ----------------------------------------------------

// Essa função recebe as requisições feitas para:
//
// /produtos
//
// Ela verifica qual método HTTP foi usado.
//
// Por exemplo:
//
// GET
// POST
func produtosHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	// r.Method contém o método HTTP utilizado.
	//
	// Por exemplo:
	//
	// GET
	// POST
	// PUT
	// DELETE
	switch r.Method {

	// Se o cliente utilizou GET:
	case http.MethodGet:

		// Chamamos a função que lista os produtos.
		listarProdutos(w)

	// Se o cliente utilizou POST:
	case http.MethodPost:

		// Chamamos a função responsável
		// por cadastrar um novo produto.
		criarProduto(w, r)

	// Se foi utilizado qualquer outro método:
	default:

		// Retornamos o status HTTP 405.
		//
		// 405 significa:
		//
		// Method Not Allowed
		//
		// Método não permitido.
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// ----------------------------------------------------
// LISTAR PRODUTOS
// ----------------------------------------------------

func listarProdutos(w http.ResponseWriter) {

	// Definimos o tipo da resposta.
	//
	// Estamos avisando para o cliente:
	//
	// "Vou devolver dados no formato JSON".
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	// Definimos o código HTTP 200.
	//
	// 200 significa que a operação ocorreu corretamente.
	w.WriteHeader(http.StatusOK)

	// Transformamos a lista produtos em JSON
	// e enviamos para o cliente.
	//
	// Exemplo:
	//
	// [
	//   {
	//      "id": 1,
	//      "nome": "Notebook",
	//      "preco": 4500
	//   }
	// ]
	json.NewEncoder(w).Encode(produtos)
}

// ----------------------------------------------------
// CRIAR UM NOVO PRODUTO
// ----------------------------------------------------

func criarProduto(
	w http.ResponseWriter,
	r *http.Request,
) {

	// Criamos uma variável vazia
	// do tipo Produto.
	//
	// Ela receberá os dados enviados
	// pelo cliente.
	var novoProduto Produto

	// r.Body contém o corpo da requisição.
	//
	// Por exemplo, o cliente pode enviar:
	//
	// {
	//   "nome": "Webcam",
	//   "preco": 299.90
	// }
	//
	// NewDecoder lê o JSON.
	//
	// Decode transforma o JSON
	// em uma struct Produto.
	err := json.NewDecoder(r.Body).Decode(&novoProduto)

	// Verificamos se aconteceu algum erro.
	//
	// Por exemplo:
	//
	// JSON escrito incorretamente.
	if err != nil {

		// Chamamos nossa função padrão
		// para enviar uma mensagem de erro.
		enviarErro(
			w,
			http.StatusBadRequest,
			"JSON inválido",
		)

		// return encerra a função.
		return
	}

	// ------------------------------------------------
	// VALIDANDO O NOME
	// ------------------------------------------------

	// Verificamos se o nome está vazio.
	if novoProduto.Nome == "" {

		// HTTP 400 significa:
		//
		// Bad Request
		//
		// O cliente enviou dados inválidos.
		enviarErro(
			w,
			http.StatusBadRequest,
			"Nome obrigatório",
		)

		return
	}

	// ------------------------------------------------
	// VALIDANDO O PREÇO
	// ------------------------------------------------

	// O preço precisa ser maior que zero.
	if novoProduto.Preco <= 0 {

		enviarErro(
			w,
			http.StatusBadRequest,
			"Preço deve ser maior que zero",
		)

		return
	}

	// ------------------------------------------------
	// GERANDO O ID
	// ------------------------------------------------

	// Criamos um ID automaticamente.
	//
	// len(produtos) retorna a quantidade de produtos.
	//
	// Exemplo:
	//
	// temos 3 produtos.
	//
	// len(produtos) = 3
	//
	// então:
	//
	// novoProduto.ID = 4
	novoProduto.ID = len(produtos) + 1

	// ------------------------------------------------
	// ADICIONANDO O PRODUTO
	// ------------------------------------------------

	// append adiciona um elemento ao final da lista.
	//
	// Portanto estamos colocando
	// novoProduto dentro da lista produtos.
	produtos = append(
		produtos,
		novoProduto,
	)

	// Informamos novamente que a resposta será JSON.
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	// Retornamos HTTP 201.
	//
	// 201 significa:
	//
	// Created
	//
	// Um novo recurso foi criado com sucesso.
	w.WriteHeader(http.StatusCreated)

	// Transformamos o produto criado em JSON
	// e enviamos para o cliente.
	json.NewEncoder(w).Encode(novoProduto)
}

// ----------------------------------------------------
// BUSCAR PRODUTO PELO ID
// ----------------------------------------------------

func produtoPorIDHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	// Essa rota aceita somente GET.
	//
	// Se o método não for GET:
	if r.Method != http.MethodGet {

		enviarErro(
			w,
			http.StatusMethodNotAllowed,
			"Utilize GET",
		)

		return
	}

	// ------------------------------------------------
	// PEGANDO O ID DA URL
	// ------------------------------------------------

	// Imagine a seguinte URL:
	//
	// /produto?id=2
	//
	// O trecho:
	//
	// id=2
	//
	// é chamado de Query Parameter.
	//
	// Query().Get("id") captura o valor
	// associado ao parâmetro id.
	idTexto := r.URL.Query().Get("id")

	// Neste momento:
	//
	// idTexto será uma string.
	//
	// Exemplo:
	//
	// "2"

	// ------------------------------------------------
	// VERIFICANDO SE O ID FOI INFORMADO
	// ------------------------------------------------

	if idTexto == "" {

		enviarErro(
			w,
			http.StatusBadRequest,
			"Informe o ID",
		)

		return
	}

	// ------------------------------------------------
	// CONVERTENDO STRING PARA INT
	// ------------------------------------------------

	// O ID veio da URL como texto.
	//
	// Precisamos transformar:
	//
	// "2"
	//
	// em:
	//
	// 2
	//
	// strconv.Atoi faz essa conversão.
	id, err := strconv.Atoi(idTexto)

	// Se alguém informar:
	//
	// /produto?id=abc
	//
	// não será possível converter "abc"
	// para um número.
	if err != nil {

		enviarErro(
			w,
			http.StatusBadRequest,
			"ID inválido",
		)

		return
	}

	// ------------------------------------------------
	// PROCURANDO O PRODUTO
	// ------------------------------------------------

	// Percorremos todos os produtos da lista.
	//
	// O range permite percorrer os elementos.
	for _, produto := range produtos {

		// O caractere _ significa:
		//
		// "não quero utilizar essa informação".
		//
		// O range poderia retornar:
		//
		// posição e produto
		//
		// Mas aqui precisamos apenas do produto.

		// Verificamos se o ID encontrado
		// é igual ao ID procurado.
		if produto.ID == id {

			// Definimos a resposta como JSON.
			w.Header().Set(
				"Content-Type",
				"application/json",
			)

			// Retornamos HTTP 200.
			w.WriteHeader(
				http.StatusOK,
			)

			// Transformamos o produto
			// encontrado em JSON.
			json.NewEncoder(w).Encode(produto)

			// Encontramos o produto,
			// então podemos encerrar a função.
			return
		}
	}

	// ------------------------------------------------
	// PRODUTO NÃO ENCONTRADO
	// ------------------------------------------------

	// Se o programa chegou até aqui,
	// significa que percorremos toda a lista
	// e nenhum produto tinha aquele ID.

	enviarErro(
		w,
		http.StatusNotFound,
		"Produto não encontrado",
	)
}

// ----------------------------------------------------
// FUNÇÃO PARA ENVIAR ERROS
// ----------------------------------------------------

// Criamos uma função específica para erros.
//
// Assim não precisamos repetir várias vezes
// o mesmo código.
func enviarErro(
	w http.ResponseWriter,
	status int,
	mensagem string,
) {

	// Toda resposta será JSON.
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	// Recebemos o status por parâmetro.
	//
	// Pode ser, por exemplo:
	//
	// 400
	// 404
	// 405
	w.WriteHeader(status)

	// Criamos um map simples.
	//
	// Por exemplo:
	//
	// map[string]string{
	//     "erro": "Produto não encontrado",
	// }
	//
	// Isso será transformado em:
	//
	// {
	//   "erro": "Produto não encontrado"
	// }
	json.NewEncoder(w).Encode(
		map[string]string{
			"erro": mensagem,
		},
	)
}
