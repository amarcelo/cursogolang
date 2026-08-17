package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	// Importa o driver que permite ao database/sql
	// conversar com o PostgreSQL.
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	// ============================================================
	// 1. CONEXÃO COM O BANCO
	// ============================================================
	//
	// Formato:
	//
	// postgres://usuario:SENHA@servidor:porta/banco
	//
	// Troque SUA_SENHA pela senha do seu PostgreSQL.
	conexao := "postgres://postgres:senha@localhost:5432/loja"

	// Abre o acesso ao PostgreSQL utilizando o driver pgx.
	db, err := sql.Open("pgx", conexao)

	if err != nil {
		fmt.Println("Erro ao abrir o banco:", err)
		return
	}

	// Quando o programa terminar, fecha o acesso ao banco.
	defer db.Close()

	// ============================================================
	// 2. TESTANDO A CONEXÃO
	// ============================================================

	// Ping verifica se o PostgreSQL está realmente acessível.
	err = db.Ping()

	if err != nil {
		fmt.Println("Erro ao conectar ao PostgreSQL:", err)
		return
	}

	fmt.Println("Conectado ao banco LOJA com sucesso!")

	// ============================================================
	// 3. PREPARANDO A LEITURA DO TECLADO
	// ============================================================

	scanner := bufio.NewScanner(os.Stdin)

	// ============================================================
	// 4. MENU PRINCIPAL
	// ============================================================

	// O programa ficará repetindo este menu até o usuário
	// escolher a opção 0.
	for {

		fmt.Println()
		fmt.Println("==============================")
		fmt.Println("    GERENCIADOR DE PRODUTOS")
		fmt.Println("==============================")
		fmt.Println("1 - Inserir novo item")
		fmt.Println("2 - Editar item por ID")
		fmt.Println("3 - Excluir item por ID")
		fmt.Println("0 - Sair")
		fmt.Println("==============================")

		fmt.Print("Escolha uma opção: ")

		// Aguarda o usuário digitar alguma coisa.
		scanner.Scan()

		opcao := strings.TrimSpace(scanner.Text())

		// Verifica qual opção foi escolhida.
		switch opcao {

		case "1":

			inserirProduto(db, scanner)

		case "2":

			editarProduto(db, scanner)

		case "3":

			excluirProduto(db, scanner)

		case "0":

			fmt.Println("Programa encerrado.")
			return

		default:

			fmt.Println("Opção inválida!")
		}
	}
}

// ================================================================
// FUNÇÃO: inserirProduto
// ================================================================
//
// Responsável por cadastrar um novo produto no banco.
func inserirProduto(db *sql.DB, scanner *bufio.Scanner) {

	fmt.Println()
	fmt.Println("--- NOVO PRODUTO ---")

	// ------------------------------------------------------------
	// Recebendo o nome
	// ------------------------------------------------------------

	fmt.Print("Nome do produto: ")

	scanner.Scan()

	nome := strings.TrimSpace(scanner.Text())

	// Não permite cadastrar um produto sem nome.
	if nome == "" {
		fmt.Println("O nome não pode ficar vazio.")
		return
	}

	// ------------------------------------------------------------
	// Recebendo o preço
	// ------------------------------------------------------------

	fmt.Print("Preço do produto: ")

	scanner.Scan()

	precoTexto := strings.TrimSpace(scanner.Text())

	// Permite que o aluno digite:
	//
	// 150.50
	//
	// ou
	//
	// 150,50
	//
	// Internamente transformamos vírgula em ponto.
	precoTexto = strings.ReplaceAll(precoTexto, ",", ".")

	// O Scanner recebe texto.
	//
	// ParseFloat transforma esse texto em um número decimal.
	preco, err := strconv.ParseFloat(precoTexto, 64)

	if err != nil {
		fmt.Println("Preço inválido.")
		return
	}

	// ------------------------------------------------------------
	// INSERT
	// ------------------------------------------------------------
	//
	// Vamos executar:
	//
	// INSERT INTO produtos
	//
	// para cadastrar o produto.
	//
	// $1 receberá o nome.
	// $2 receberá o preço.
	//
	// RETURNING id pede ao PostgreSQL que devolva o ID
	// criado automaticamente.

	var id int

	err = db.QueryRow(
		`
		INSERT INTO produtos (nome, preco)
		VALUES ($1, $2)
		RETURNING id
		`,
		nome,
		preco,
	).Scan(&id)

	if err != nil {
		fmt.Println("Erro ao cadastrar produto:", err)
		return
	}

	fmt.Println()
	fmt.Println("Produto cadastrado com sucesso!")
	fmt.Println("ID criado:", id)
}

// ================================================================
// FUNÇÃO: editarProduto
// ================================================================
//
// Permite modificar um produto existente.
//
// Primeiro perguntamos o ID.
// Depois procuramos o produto.
// Finalmente alteramos nome e preço.
func editarProduto(db *sql.DB, scanner *bufio.Scanner) {

	fmt.Println()
	fmt.Println("--- EDITAR PRODUTO ---")

	// ------------------------------------------------------------
	// Recebendo o ID
	// ------------------------------------------------------------

	fmt.Print("Digite o ID do produto: ")

	scanner.Scan()

	idTexto := strings.TrimSpace(scanner.Text())

	// Converte o ID digitado para número inteiro.
	id, err := strconv.Atoi(idTexto)

	if err != nil {
		fmt.Println("ID inválido.")
		return
	}

	// ------------------------------------------------------------
	// PROCURANDO O PRODUTO
	// ------------------------------------------------------------
	//
	// Antes de modificar, verificamos se o produto existe.

	var nomeAtual string
	var precoAtual float64

	err = db.QueryRow(
		`
		SELECT nome, preco
		FROM produtos
		WHERE id = $1
		`,
		id,
	).Scan(
		&nomeAtual,
		&precoAtual,
	)

	// ErrNoRows significa que o SELECT não encontrou
	// nenhum registro com aquele ID.
	if err == sql.ErrNoRows {
		fmt.Println("Produto não encontrado.")
		return
	}

	if err != nil {
		fmt.Println("Erro ao procurar produto:", err)
		return
	}

	// ------------------------------------------------------------
	// MOSTRANDO O PRODUTO ENCONTRADO
	// ------------------------------------------------------------

	fmt.Println()
	fmt.Println("Produto encontrado:")
	fmt.Println("ID:", id)
	fmt.Println("Nome:", nomeAtual)
	fmt.Printf("Preço: %.2f\n", precoAtual)

	// ------------------------------------------------------------
	// RECEBENDO O NOVO NOME
	// ------------------------------------------------------------

	fmt.Print("Novo nome: ")

	scanner.Scan()

	novoNome := strings.TrimSpace(scanner.Text())

	if novoNome == "" {
		fmt.Println("O nome não pode ficar vazio.")
		return
	}

	// ------------------------------------------------------------
	// RECEBENDO O NOVO PREÇO
	// ------------------------------------------------------------

	fmt.Print("Novo preço: ")

	scanner.Scan()

	precoTexto := strings.TrimSpace(scanner.Text())

	precoTexto = strings.ReplaceAll(precoTexto, ",", ".")

	novoPreco, err := strconv.ParseFloat(precoTexto, 64)

	if err != nil {
		fmt.Println("Preço inválido.")
		return
	}

	// ------------------------------------------------------------
	// UPDATE
	// ------------------------------------------------------------
	//
	// Atualizamos nome e preço do produto cujo ID foi informado.

	resultado, err := db.Exec(
		`
		UPDATE produtos
		SET nome = $1,
		    preco = $2
		WHERE id = $3
		`,
		novoNome,
		novoPreco,
		id,
	)

	if err != nil {
		fmt.Println("Erro ao atualizar produto:", err)
		return
	}

	// Descobrimos quantos registros foram modificados.
	linhas, err := resultado.RowsAffected()

	if err != nil {
		fmt.Println("Erro ao verificar atualização:", err)
		return
	}

	if linhas == 0 {
		fmt.Println("Nenhum produto foi atualizado.")
		return
	}

	fmt.Println("Produto atualizado com sucesso!")
}

// ================================================================
// FUNÇÃO: excluirProduto
// ================================================================
//
// Permite excluir um produto utilizando seu ID.
func excluirProduto(db *sql.DB, scanner *bufio.Scanner) {

	fmt.Println()
	fmt.Println("--- EXCLUIR PRODUTO ---")

	// ------------------------------------------------------------
	// RECEBENDO O ID
	// ------------------------------------------------------------

	fmt.Print("Digite o ID do produto: ")

	scanner.Scan()

	idTexto := strings.TrimSpace(scanner.Text())

	id, err := strconv.Atoi(idTexto)

	if err != nil {
		fmt.Println("ID inválido.")
		return
	}

	// ------------------------------------------------------------
	// PROCURANDO O PRODUTO
	// ------------------------------------------------------------
	//
	// Antes de apagar, verificamos se ele realmente existe.

	var nome string

	err = db.QueryRow(
		`
		SELECT nome
		FROM produtos
		WHERE id = $1
		`,
		id,
	).Scan(&nome)

	if err == sql.ErrNoRows {
		fmt.Println("Produto não encontrado.")
		return
	}

	if err != nil {
		fmt.Println("Erro ao procurar produto:", err)
		return
	}

	// Mostramos qual produto será excluído.
	fmt.Println()
	fmt.Println("Produto encontrado:", nome)

	// ------------------------------------------------------------
	// CONFIRMAÇÃO
	// ------------------------------------------------------------

	fmt.Print("Tem certeza que deseja excluir? (s/n): ")

	scanner.Scan()

	resposta := strings.ToLower(
		strings.TrimSpace(scanner.Text()),
	)

	// Somente "s" confirma a exclusão.
	if resposta != "s" {
		fmt.Println("Exclusão cancelada.")
		return
	}

	// ------------------------------------------------------------
	// DELETE
	// ------------------------------------------------------------

	resultado, err := db.Exec(
		`
		DELETE FROM produtos
		WHERE id = $1
		`,
		id,
	)

	if err != nil {
		fmt.Println("Erro ao excluir produto:", err)
		return
	}

	// Verifica quantas linhas foram apagadas.
	linhas, err := resultado.RowsAffected()

	if err != nil {
		fmt.Println("Erro ao verificar exclusão:", err)
		return
	}

	if linhas == 0 {
		fmt.Println("Nenhum produto foi excluído.")
		return
	}

	fmt.Println("Produto excluído com sucesso!")
}
