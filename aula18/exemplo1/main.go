package main

// Importamos os pacotes que vamos utilizar no programa.
import (
	"encoding/json" // Trabalhar com JSON
	"fmt"           // Mostrar mensagens no terminal
	"log"           // Criar logs
	"net/http"      // Criar o servidor HTTP e trabalhar com requisições
	"strconv"       // Converter texto para número
)

// ---------------------------------------------------------
// STRUCT DE PRODUTO
// ---------------------------------------------------------

// Product representa a estrutura de um produto.
//
// Cada produto terá:
//
// ID
// Nome
// Preço
//
// As partes entre `json:"..."` dizem qual será o nome
// do campo quando o produto for transformado em JSON.
type Product struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

// ---------------------------------------------------------
// LISTA DE PRODUTOS
// ---------------------------------------------------------

// Aqui criamos uma lista de produtos.
//
// Neste exemplo NÃO estamos utilizando banco de dados.
//
// Os produtos ficarão armazenados apenas na memória.
//
// Isso significa que, se o servidor for desligado,
// os produtos cadastrados durante a execução serão perdidos.
var produtos = []Product{
	{ID: 1, Nome: "Notebook", Preco: 3500.00},
	{ID: 2, Nome: "Mouse", Preco: 120.00},
	{ID: 3, Nome: "Teclado", Preco: 250.00},
}

// ---------------------------------------------------------
// FUNÇÃO PRINCIPAL
// ---------------------------------------------------------

func main() {

	// Primeiro transformamos a função produtosRouter
	// em um Handler HTTP.
	//
	// Um Handler é uma função responsável por atender
	// uma requisição HTTP.
	produtosHandler := http.HandlerFunc(produtosRouter)

	// Aqui registramos a rota /produtos.
	//
	// Antes de chegar ao produtosHandler,
	// a requisição passará pelo middleware de logging.
	//
	// Portanto o fluxo será:
	//
	// Cliente
	//    ↓
	// logging
	//    ↓
	// produtosRouter
	http.Handle(
		"/produtos",
		logging(produtosHandler),
	)

	// Apenas mostramos mensagens no terminal.
	fmt.Println("Servidor iniciado em:")
	fmt.Println("http://localhost:8080")

	// Inicia o servidor HTTP.
	//
	// :8080 significa que o servidor ficará
	// ouvindo requisições na porta 8080.
	//
	// O segundo parâmetro nil significa que estamos
	// utilizando o roteador padrão do pacote net/http.
	err := http.ListenAndServe(":8080", nil)

	// Caso aconteça algum erro ao iniciar o servidor,
	// mostramos o erro e encerramos o programa.
	if err != nil {
		log.Fatal(err)
	}
}

// ---------------------------------------------------------
// MIDDLEWARE DE LOGGING
// ---------------------------------------------------------

// logging é nosso primeiro Middleware.
//
// Um Middleware recebe um Handler chamado "next".
//
// Esse "next" representa o próximo passo
// que será executado depois do Middleware.
func logging(next http.Handler) http.Handler {

	// Retornamos um novo Handler.
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			// Mostramos no terminal informações
			// sobre a requisição recebida.
			//
			// r.Method mostra o método HTTP:
			//
			// GET
			// POST
			// PATCH
			// DELETE
			//
			// r.URL.Path mostra a rota acessada.
			log.Printf(
				"Método: %s | Rota: %s",
				r.Method,
				r.URL.Path,
			)

			// Esta linha significa:
			//
			// "O Middleware terminou seu trabalho.
			// Pode continuar para o próximo Handler."
			next.ServeHTTP(w, r)
		},
	)
}

// ---------------------------------------------------------
// MIDDLEWARE DE AUTENTICAÇÃO
// ---------------------------------------------------------

// Este Middleware verifica se o usuário
// informou login e senha corretos.
//
// Neste exemplo vamos utilizar Basic Authentication.
func autenticacao(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			// BasicAuth tenta recuperar o usuário e a senha
			// enviados na requisição.
			//
			// usuario recebe o nome do usuário.
			// senha recebe a senha.
			// ok informa se a autenticação foi enviada.
			usuario, senha, ok := r.BasicAuth()

			// Verificamos três situações:
			//
			// 1. Não foi enviada autenticação
			// 2. O usuário não é admin
			// 3. A senha não é 1234
			if !ok ||
				usuario != "admin" ||
				senha != "1234" {

				// Este cabeçalho informa ao cliente
				// que esta rota utiliza Basic Authentication.
				w.Header().Set(
					"WWW-Authenticate",
					`Basic realm="API de Produtos"`,
				)

				// Enviamos uma mensagem de erro.
				//
				// StatusUnauthorized representa:
				//
				// HTTP 401
				//
				// Significa "Não autorizado".
				http.Error(
					w,
					"Acesso não autorizado",
					http.StatusUnauthorized,
				)

				// IMPORTANTE:
				//
				// return interrompe a função.
				//
				// Como a autenticação falhou,
				// não deixamos a requisição continuar.
				return
			}

			// Se chegou aqui, usuário e senha estão corretos.
			//
			// Então permitimos que a requisição continue.
			next.ServeHTTP(w, r)
		},
	)
}

// ---------------------------------------------------------
// ROTEADOR PRINCIPAL
// ---------------------------------------------------------

// produtosRouter recebe todas as requisições
// feitas para:
//
// /produtos
//
// Depois verifica qual método HTTP foi utilizado.
func produtosRouter(
	w http.ResponseWriter,
	r *http.Request,
) {

	// switch permite tomar decisões
	// com base no método HTTP.
	switch r.Method {

	// -------------------------------------------------
	// GET
	// -------------------------------------------------

	case http.MethodGet:

		// GET será utilizado para consultar produtos.
		getProdutos(w, r)

	// -------------------------------------------------
	// POST
	// -------------------------------------------------

	case http.MethodPost:

		// POST cria produtos.
		//
		// Como estamos alterando dados,
		// exigimos autenticação.
		//
		// Primeiro transformamos postProduto em Handler.
		//
		// Depois colocamos o Middleware de autenticação.
		//
		// Finalmente executamos tudo com ServeHTTP.
		autenticacao(
			http.HandlerFunc(postProduto),
		).ServeHTTP(w, r)

	// -------------------------------------------------
	// PATCH
	// -------------------------------------------------

	case http.MethodPatch:

		// PATCH altera parcialmente um produto.
		//
		// Também exigimos autenticação.
		autenticacao(
			http.HandlerFunc(patchProduto),
		).ServeHTTP(w, r)

	// -------------------------------------------------
	// DELETE
	// -------------------------------------------------

	case http.MethodDelete:

		// DELETE remove um produto.
		//
		// Também exigimos autenticação.
		autenticacao(
			http.HandlerFunc(deleteProduto),
		).ServeHTTP(w, r)

	// -------------------------------------------------
	// QUALQUER OUTRO MÉTODO
	// -------------------------------------------------

	default:

		// Caso alguém tente usar outro método,
		// como PUT, OPTIONS etc.,
		// respondemos com erro 405.
		http.Error(
			w,
			"Método não permitido",
			http.StatusMethodNotAllowed,
		)
	}
}

// ---------------------------------------------------------
// GET
// ---------------------------------------------------------

// getProdutos permite:
//
// GET /produtos
//
// ou:
//
// GET /produtos?id=1
func getProdutos(
	w http.ResponseWriter,
	r *http.Request,
) {

	// Pegamos o valor do parâmetro "id".
	//
	// Exemplo:
	//
	// /produtos?id=2
	//
	// Nesse caso idTexto será "2".
	idTexto := r.URL.Query().Get("id")

	// Se o parâmetro id NÃO foi enviado...
	if idTexto == "" {

		// ...retornamos todos os produtos.
		responderJSON(
			w,
			http.StatusOK,
			produtos,
		)

		// Encerramos a função.
		return
	}

	// O ID chega pela URL como texto.
	//
	// Precisamos convertê-lo para número inteiro.
	//
	// strconv.Atoi faz essa conversão.
	id, err := strconv.Atoi(idTexto)

	// Caso o usuário escreva algo inválido:
	//
	// /produtos?id=abc
	//
	// a conversão dará erro.
	if err != nil {

		http.Error(
			w,
			"ID inválido",
			http.StatusBadRequest,
		)

		return
	}

	// Agora percorremos todos os produtos.
	for _, produto := range produtos {

		// Comparamos o ID do produto
		// com o ID solicitado.
		if produto.ID == id {

			// Encontramos o produto.
			//
			// Retornamos o produto em JSON.
			responderJSON(
				w,
				http.StatusOK,
				produto,
			)

			return
		}
	}

	// Se percorremos todos os produtos
	// e não encontramos o ID,
	// retornamos HTTP 404.
	http.Error(
		w,
		"Produto não encontrado",
		http.StatusNotFound,
	)
}

// ---------------------------------------------------------
// POST
// ---------------------------------------------------------

// postProduto cria um novo produto.
//
// Esperamos receber um JSON parecido com:
//
//	{
//	    "nome": "Webcam",
//	    "preco": 299.90
//	}
func postProduto(
	w http.ResponseWriter,
	r *http.Request,
) {

	// Criamos uma variável vazia.
	//
	// Ela receberá os dados enviados pelo cliente.
	var novoProduto Product

	// NewDecoder lê o corpo da requisição.
	//
	// Decode transforma o JSON recebido
	// em uma struct Go.
	//
	// Observe o &novoProduto.
	//
	// Estamos passando o endereço da variável
	// para que Decode possa modificar seu conteúdo.
	err := json.NewDecoder(
		r.Body,
	).Decode(&novoProduto)

	// Se o JSON estiver errado...
	if err != nil {

		http.Error(
			w,
			"JSON inválido",
			http.StatusBadRequest,
		)

		return
	}

	// Criamos automaticamente um ID.
	//
	// Como é um exemplo simples,
	// usamos o tamanho da lista + 1.
	//
	// Em uma aplicação real,
	// normalmente o banco de dados faria isso.
	novoProduto.ID = len(produtos) + 1

	// append adiciona o novo produto
	// ao final da lista.
	produtos = append(
		produtos,
		novoProduto,
	)

	// Retornamos o produto criado.
	//
	// StatusCreated representa HTTP 201.
	responderJSON(
		w,
		http.StatusCreated,
		novoProduto,
	)
}

// ---------------------------------------------------------
// PATCH
// ---------------------------------------------------------

// PATCH é utilizado para alterar
// apenas parte de um produto.
//
// Exemplo:
//
// PATCH /produtos?id=2
//
// JSON:
//
//	{
//	    "preco": 150
//	}
func patchProduto(
	w http.ResponseWriter,
	r *http.Request,
) {

	// Pegamos o ID enviado na URL.
	idTexto := r.URL.Query().Get("id")

	// Transformamos o ID de texto para número.
	id, err := strconv.Atoi(idTexto)

	if err != nil {

		http.Error(
			w,
			"ID inválido",
			http.StatusBadRequest,
		)

		return
	}

	// Criamos uma estrutura temporária
	// para receber os campos enviados.
	//
	// Aqui estamos utilizando ponteiros:
	//
	// *string
	// *float64
	//
	// Isso permite descobrir se o campo
	// realmente foi enviado no JSON.
	var dados struct {
		Nome  *string  `json:"nome"`
		Preco *float64 `json:"preco"`
	}

	// Convertemos o JSON recebido
	// para nossa estrutura temporária.
	err = json.NewDecoder(
		r.Body,
	).Decode(&dados)

	if err != nil {

		http.Error(
			w,
			"JSON inválido",
			http.StatusBadRequest,
		)

		return
	}

	// Percorremos a lista de produtos.
	//
	// Aqui utilizamos:
	//
	// for i := range produtos
	//
	// porque precisamos do índice
	// para alterar o produto dentro da lista.
	for i := range produtos {

		// Encontramos o produto?
		if produtos[i].ID == id {

			// Verificamos se o nome foi enviado.
			if dados.Nome != nil {

				// *dados.Nome significa:
				//
				// acessar o valor armazenado
				// naquele endereço de memória.
				produtos[i].Nome = *dados.Nome
			}

			// Verificamos se o preço foi enviado.
			if dados.Preco != nil {

				produtos[i].Preco = *dados.Preco
			}

			// Retornamos o produto atualizado.
			responderJSON(
				w,
				http.StatusOK,
				produtos[i],
			)

			return
		}
	}

	// Produto não encontrado.
	http.Error(
		w,
		"Produto não encontrado",
		http.StatusNotFound,
	)
}

// ---------------------------------------------------------
// DELETE
// ---------------------------------------------------------

// deleteProduto remove um produto.
//
// Exemplo:
//
// DELETE /produtos?id=3
func deleteProduto(
	w http.ResponseWriter,
	r *http.Request,
) {

	// Pegamos o ID.
	idTexto := r.URL.Query().Get("id")

	// Convertemos para inteiro.
	id, err := strconv.Atoi(idTexto)

	if err != nil {

		http.Error(
			w,
			"ID inválido",
			http.StatusBadRequest,
		)

		return
	}

	// Percorremos os produtos.
	//
	// Aqui precisamos do índice "i"
	// porque vamos remover um elemento da lista.
	for i, produto := range produtos {

		// Encontramos o produto?
		if produto.ID == id {

			// Aqui removemos o produto da lista.
			//
			// Exemplo:
			//
			// produtos:
			//
			// [Notebook, Mouse, Teclado]
			//
			// Se queremos remover Mouse,
			// juntamos:
			//
			// [Notebook]
			//
			// com
			//
			// [Teclado]
			//
			// O "..." permite adicionar
			// todos os elementos da segunda parte.
			produtos = append(
				produtos[:i],
				produtos[i+1:]...,
			)

			// HTTP 204 significa:
			//
			// operação realizada com sucesso,
			// mas não existe conteúdo para retornar.
			w.WriteHeader(
				http.StatusNoContent,
			)

			return
		}
	}

	// Se não encontrou o produto:
	http.Error(
		w,
		"Produto não encontrado",
		http.StatusNotFound,
	)
}

// ---------------------------------------------------------
// FUNÇÃO AUXILIAR PARA JSON
// ---------------------------------------------------------

// responderJSON evita repetir o mesmo código
// em várias partes da aplicação.
//
// Ela recebe:
//
// w      -> objeto usado para criar a resposta
// status -> código HTTP
// dados  -> informação que será convertida para JSON
func responderJSON(
	w http.ResponseWriter,
	status int,
	dados interface{},
) {

	// Informamos ao cliente que a resposta
	// enviada será no formato JSON.
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	// Define o código HTTP.
	//
	// Exemplos:
	//
	// 200 OK
	// 201 Created
	// 404 Not Found
	w.WriteHeader(status)

	// NewEncoder cria um codificador JSON.
	//
	// Encode transforma os dados Go em JSON
	// e envia diretamente para o cliente.
	json.NewEncoder(w).Encode(dados)
}
