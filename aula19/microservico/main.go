package main

import (
	"fmt"
	"log"
	"microservico/controller"
	"net/http"
)

func main() {

	http.HandleFunc(
		"/produtos",
		controller.ProdutosHandler,
	)

	http.HandleFunc(
		"/produtos/",
		controller.ProdutoHandler,
	)

	fmt.Println(
		"Servidor rodando em http://localhost:8080",
	)

	err := http.ListenAndServe(
		":8080",
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}
}
