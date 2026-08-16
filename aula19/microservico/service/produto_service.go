package service

import "microservico/model"

// produtos funciona como nosso banco de dados temporário.
var produtos = []model.Produto{
	{
		ID:    1,
		Nome:  "Notebook",
		Preco: 3500,
	},
	{
		ID:    2,
		Nome:  "Mouse",
		Preco: 100,
	},
}

// ListarProdutos retorna todos os produtos.
func ListarProdutos() []model.Produto {
	return produtos
}

// BuscarProduto procura um produto pelo ID.
func BuscarProduto(id int) (model.Produto, bool) {

	for _, produto := range produtos {

		if produto.ID == id {
			return produto, true
		}
	}

	return model.Produto{}, false
}

// CriarProduto adiciona um novo produto.
func CriarProduto(produto model.Produto) model.Produto {

	produto.ID = len(produtos) + 1

	produtos = append(produtos, produto)

	return produto
}
