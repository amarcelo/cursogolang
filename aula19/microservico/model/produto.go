package model

// Produto representa um produto dentro da aplicação.
type Produto struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}
