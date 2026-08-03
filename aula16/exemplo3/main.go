package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Estrutura que representa um produto.
type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/products", productsHandler)

	fmt.Println("Servidor rodando em :8080")
	http.ListenAndServe(":8080", nil)
}

// Página inicial.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Online"))
}

// Retorna um único produto pelo ID.
func productsHandler(w http.ResponseWriter, r *http.Request) {

	// Lista de produtos.
	products := []Product{
		{ID: 1, Name: "Notebook"},
		{ID: 2, Name: "Mouse"},
		{ID: 3, Name: "Teclado"},
	}

	// Obtém o parâmetro "id" da URL.
	idStr := r.URL.Query().Get("id")

	// Converte o ID para inteiro.
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Procura o produto.
	for _, product := range products {
		if product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	// Caso não encontre.
	http.Error(w, "Produto não encontrado", http.StatusNotFound)
}
