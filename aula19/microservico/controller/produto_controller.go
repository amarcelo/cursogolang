package controller

import (
	"encoding/json"
	"microservico/model"
	"microservico/service"
	"net/http"
	"strconv"
	"strings"
)

// ProdutosHandler controla:
//
// GET  /produtos
// POST /produtos
func ProdutosHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {

	case http.MethodGet:

		produtos := service.ListarProdutos()

		json.NewEncoder(w).Encode(produtos)

	case http.MethodPost:

		var produto model.Produto

		err := json.NewDecoder(r.Body).Decode(&produto)

		if err != nil {

			http.Error(
				w,
				"JSON inválido",
				http.StatusBadRequest,
			)

			return
		}

		produto = service.CriarProduto(produto)

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(produto)

	default:

		http.Error(
			w,
			"Método não permitido",
			http.StatusMethodNotAllowed,
		)
	}
}

// ProdutoHandler controla:
//
// GET /produtos/1
// GET /produtos/2
func ProdutoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		http.Error(
			w,
			"Método não permitido",
			http.StatusMethodNotAllowed,
		)

		return
	}

	idTexto := strings.TrimPrefix(
		r.URL.Path,
		"/produtos/",
	)

	id, err := strconv.Atoi(idTexto)

	if err != nil {

		http.Error(
			w,
			"ID inválido",
			http.StatusBadRequest,
		)

		return
	}

	produto, encontrado := service.BuscarProduto(id)

	if !encontrado {

		http.Error(
			w,
			"Produto não encontrado",
			http.StatusNotFound,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(produto)
}
