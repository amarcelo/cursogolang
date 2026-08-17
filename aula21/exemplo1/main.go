package main

import (
	// Pacote padrão do Go para trabalhar com bancos de dados SQL
	"database/sql"

	// Pacote usado para mostrar informações na tela
	"fmt"

	// Driver que permite ao Go conversar com o PostgreSQL
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	// ============================================================
	// 1. CONFIGURAÇÃO DA CONEXÃO
	// ============================================================

	// String contendo os dados necessários para acessar o PostgreSQL
	//
	// postgres://usuario:senha@servidor:porta/banco
	//
	// Troque "senha" pela senha real do seu PostgreSQL.
	conexao := "postgres://postgres:senha@localhost:5432/loja"

	// ============================================================
	// 2. ABRINDO O ACESSO AO BANCO
	// ============================================================

	// Prepara o acesso ao PostgreSQL usando o driver pgx
	db, err := sql.Open("pgx", conexao)

	// Verifica se ocorreu algum erro
	if err != nil {
		panic(err)
	}

	// Quando o programa terminar, fecha o acesso ao banco
	defer db.Close()

	// ============================================================
	// 3. TESTANDO A CONEXÃO
	// ============================================================

	// Ping() verifica se o PostgreSQL está realmente respondendo
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Conectado ao PostgreSQL!")

	// ============================================================
	// 4. CADASTRANDO UM PRODUTO
	// ============================================================

	// Exec() executa um comando SQL.
	//
	// Neste caso estamos cadastrando:
	//
	// nome  = Monitor
	// preco = 1200.00
	//
	// $1 recebe "Monitor"
	// $2 recebe 1200.00
	_, err = db.Exec(
		`INSERT INTO produtos
		(nome, preco)
		VALUES ($1, $2)`,
		"Notebook",
		3600.00,
	)

	// Verifica se ocorreu algum erro durante o INSERT
	if err != nil {
		panic(err)
	}

	fmt.Println("Produto cadastrado!")

	// ============================================================
	// 5. CONSULTANDO OS PRODUTOS
	// ============================================================

	// Query() executa uma consulta que retorna dados.
	//
	// Estamos buscando:
	//
	// id
	// nome
	// preco
	rows, err := db.Query(
		"SELECT id, nome, preco FROM produtos",
	)

	if err != nil {
		panic(err)
	}

	// Fecha o resultado da consulta quando o programa terminar
	defer rows.Close()

	fmt.Println("\nLISTA DE PRODUTOS")

	// ============================================================
	// 6. PERCORRENDO OS PRODUTOS
	// ============================================================

	// Next() avança para o próximo produto encontrado
	for rows.Next() {

		// Variáveis que receberão os dados do banco
		var id int
		var nome string
		var preco float64

		// ========================================================
		// 7. COPIANDO OS DADOS PARA AS VARIÁVEIS
		// ========================================================

		// O SELECT retorna:
		//
		// id, nome, preco
		//
		// Portanto o Scan deve seguir a mesma ordem:
		//
		// &id, &nome, &preco
		err := rows.Scan(
			&id,
			&nome,
			&preco,
		)

		if err != nil {
			panic(err)
		}

		// ========================================================
		// 8. MOSTRANDO O PRODUTO
		// ========================================================

		fmt.Println("-------------------")
		fmt.Println("ID:", id)
		fmt.Println("Nome:", nome)
		fmt.Println("Preço:", preco)
	}
}
