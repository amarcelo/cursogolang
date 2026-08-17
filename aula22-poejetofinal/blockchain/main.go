package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// ============================================================
// CONFIGURAÇÕES GERAIS
// ============================================================

// Nome do arquivo onde a blockchain ficará armazenada.
const ledgerFile = "ledger.json"

// Dificuldade da mineração.
//
// 4 significa que o hash precisa começar com:
//
// 0000
//
// Quanto maior esse número, mais difícil será minerar.
const difficulty = 4

// ============================================================
// ESTRUTURA DE UMA TRANSAÇÃO
// ============================================================

// Transaction representa uma transferência de valores.
//
// Exemplo:
//
// João -> Maria -> R$ 500
type Transaction struct {

	// Quem está enviando.
	From string `json:"from"`

	// Quem está recebendo.
	To string `json:"to"`

	// Valor da transação.
	Amount float64 `json:"amount"`
}

// ============================================================
// ESTRUTURA DE UM BLOCO
// ============================================================

type Block struct {

	// Número sequencial do bloco.
	Index int `json:"index"`

	// Data e hora em que o bloco foi criado.
	Timestamp string `json:"timestamp"`

	// Lista de transações armazenadas neste bloco.
	Transactions []Transaction `json:"transactions"`

	// Nome da pessoa que minerou o bloco.
	Miner string `json:"miner"`

	// Hash do bloco anterior.
	PreviousHash string `json:"previous_hash"`

	// Hash do próprio bloco.
	Hash string `json:"hash"`

	// Número utilizado no processo de Proof of Work.
	Nonce int `json:"nonce"`

	// Quantidade de tentativas realizadas durante a mineração.
	Attempts int `json:"attempts"`

	// Tempo gasto para minerar o bloco.
	MiningTime string `json:"mining_time"`
}

// ============================================================
// FUNÇÃO PARA CALCULAR O HASH
// ============================================================

// calculateHash recebe um bloco e calcula sua impressão
// digital usando SHA-256.
func calculateHash(block Block) string {

	// Primeiro precisamos transformar as transações em JSON.
	//
	// Dessa forma elas também participam do cálculo do hash.
	transactionsJSON, err := json.Marshal(block.Transactions)

	if err != nil {
		return ""
	}

	// Criamos uma grande string contendo os dados importantes
	// do bloco.
	record := strconv.Itoa(block.Index) +
		block.Timestamp +
		string(transactionsJSON) +
		block.Miner +
		block.PreviousHash +
		strconv.Itoa(block.Nonce)

	// Calculamos SHA-256.
	hash := sha256.Sum256([]byte(record))

	// O resultado do SHA-256 é convertido para hexadecimal.
	return hex.EncodeToString(hash[:])
}

// ============================================================
// PROOF OF WORK
// ============================================================

// mineBlock executa o processo de mineração.
//
// Recebemos um ponteiro:
//
// *Block
//
// porque queremos modificar diretamente o bloco original.
func mineBlock(block *Block) {

	// Cria o alvo da mineração.
	//
	// Se difficulty = 4:
	//
	// target = "0000"
	target := strings.Repeat("0", difficulty)

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("INICIANDO MINERAÇÃO")
	fmt.Println("========================================")
	fmt.Println("Bloco:", block.Index)
	fmt.Println("Minerador:", block.Miner)
	fmt.Println("Dificuldade:", difficulty)
	fmt.Println("Alvo:", target)
	fmt.Println()

	// Marcamos o horário em que a mineração começou.
	start := time.Now()

	// Começamos o nonce com zero.
	block.Nonce = 0

	// Contador de tentativas.
	block.Attempts = 0

	// Loop infinito.
	//
	// Ele somente termina quando encontrarmos um hash válido.
	for {

		// Uma nova tentativa está sendo realizada.
		block.Attempts++

		// Calculamos o hash atual.
		block.Hash = calculateHash(*block)

		// Verificamos se o hash começa com a quantidade
		// necessária de zeros.
		if strings.HasPrefix(block.Hash, target) {

			// Encontramos!
			break
		}

		// Caso contrário, alteramos o nonce.
		block.Nonce++
	}

	// Calculamos quanto tempo foi gasto.
	duration := time.Since(start)

	// Guardamos o tempo dentro do bloco.
	block.MiningTime = duration.String()

	fmt.Println("BLOCO MINERADO COM SUCESSO!")
	fmt.Println()
	fmt.Println("Nonce encontrado:", block.Nonce)
	fmt.Println("Tentativas:", block.Attempts)
	fmt.Println("Tempo de mineração:", block.MiningTime)
	fmt.Println("Hash:", block.Hash)
	fmt.Println("========================================")
}

// ============================================================
// GENESIS BLOCK
// ============================================================

// createGenesisBlock cria o primeiro bloco da blockchain.
func createGenesisBlock() Block {

	fmt.Println()
	fmt.Println("Criando Genesis Block...")

	// Criamos uma transação especial apenas para identificar
	// o Genesis Block.
	genesisTransaction := Transaction{
		From:   "Sistema",
		To:     "Blockchain",
		Amount: 0,
	}

	block := Block{
		Index: 0,

		Timestamp: time.Now().Format(time.RFC3339),

		Transactions: []Transaction{
			genesisTransaction,
		},

		Miner: "Sistema",

		// Como não existe bloco anterior, usamos "0".
		PreviousHash: "0",

		Nonce: 0,
	}

	// Até o Genesis Block precisa passar pelo Proof of Work.
	mineBlock(&block)

	return block
}

// ============================================================
// CRIAR NOVO BLOCO
// ============================================================

// createBlock cria um novo bloco.
//
// Recebe:
//
// transactions  -> transações que serão colocadas no bloco
// miner         -> nome do minerador
// previousBlock -> bloco anterior
func createBlock(
	transactions []Transaction,
	miner string,
	previousBlock Block,
) Block {

	block := Block{

		// O índice será o índice anterior + 1.
		Index: previousBlock.Index + 1,

		// Data e hora atual.
		Timestamp: time.Now().Format(time.RFC3339),

		// Transações que estavam aguardando mineração.
		Transactions: transactions,

		// Nome do minerador.
		Miner: miner,

		// Ligamos este bloco ao bloco anterior.
		PreviousHash: previousBlock.Hash,

		// A mineração começa no nonce zero.
		Nonce: 0,
	}

	// Executamos o Proof of Work.
	mineBlock(&block)

	return block
}

// ============================================================
// SALVAR BLOCKCHAIN
// ============================================================

// saveBlockchain grava toda a blockchain no arquivo ledger.json.
func saveBlockchain(blockchain []Block) error {

	// Converte as structs Go para JSON formatado.
	data, err := json.MarshalIndent(blockchain, "", "    ")

	if err != nil {
		return err
	}

	// Grava o arquivo.
	err = os.WriteFile(
		ledgerFile,
		data,
		0644,
	)

	if err != nil {
		return err
	}

	return nil
}

// ============================================================
// CARREGAR BLOCKCHAIN
// ============================================================

// loadBlockchain lê o ledger.json e reconstrói
// nossa blockchain na memória.
func loadBlockchain() ([]Block, error) {

	// Lê o conteúdo do arquivo.
	data, err := os.ReadFile(ledgerFile)

	if err != nil {
		return nil, err
	}

	// Criamos uma variável que receberá os blocos.
	var blockchain []Block

	// Converte JSON para estruturas Go.
	err = json.Unmarshal(data, &blockchain)

	if err != nil {
		return nil, err
	}

	return blockchain, nil
}

// ============================================================
// VALIDAR BLOCKCHAIN
// ============================================================

// validateBlockchain verifica toda a integridade da blockchain.
func validateBlockchain(blockchain []Block) bool {

	// Uma blockchain vazia não é válida.
	if len(blockchain) == 0 {
		return false
	}

	// Primeiro verificamos o Genesis Block.
	genesis := blockchain[0]

	// Recalculamos seu hash.
	if calculateHash(genesis) != genesis.Hash {
		fmt.Println("Erro no Genesis Block.")
		return false
	}

	// Verificamos também o Proof of Work do Genesis.
	target := strings.Repeat("0", difficulty)

	if !strings.HasPrefix(genesis.Hash, target) {
		fmt.Println("Proof of Work inválido no Genesis Block.")
		return false
	}

	// Agora percorremos os demais blocos.
	for i := 1; i < len(blockchain); i++ {

		currentBlock := blockchain[i]

		previousBlock := blockchain[i-1]

		// ----------------------------------------------------
		// VERIFICAÇÃO 1
		// Índices precisam ser sequenciais.
		// ----------------------------------------------------

		if currentBlock.Index != previousBlock.Index+1 {

			fmt.Println("Erro de índice no bloco:", currentBlock.Index)

			return false
		}

		// ----------------------------------------------------
		// VERIFICAÇÃO 2
		// PreviousHash precisa apontar para o bloco anterior.
		// ----------------------------------------------------

		if currentBlock.PreviousHash != previousBlock.Hash {

			fmt.Println("Ligação quebrada no bloco:", currentBlock.Index)

			return false
		}

		// ----------------------------------------------------
		// VERIFICAÇÃO 3
		// Recalculamos o hash.
		// ----------------------------------------------------

		calculatedHash := calculateHash(currentBlock)

		if calculatedHash != currentBlock.Hash {

			fmt.Println("Hash inválido no bloco:", currentBlock.Index)

			return false
		}

		// ----------------------------------------------------
		// VERIFICAÇÃO 4
		// Verificamos o Proof of Work.
		// ----------------------------------------------------

		if !strings.HasPrefix(currentBlock.Hash, target) {

			fmt.Println(
				"Proof of Work inválido no bloco:",
				currentBlock.Index,
			)

			return false
		}
	}

	// Se passamos por todos os testes,
	// a blockchain está válida.
	return true
}

// ============================================================
// MOSTRAR UMA TRANSAÇÃO
// ============================================================

func showTransaction(transaction Transaction) {

	fmt.Println("    Origem:", transaction.From)
	fmt.Println("    Destino:", transaction.To)

	fmt.Printf(
		"    Valor: R$ %.2f\n",
		transaction.Amount,
	)
}

// ============================================================
// MOSTRAR BLOCKCHAIN
// ============================================================

func showBlockchain(blockchain []Block) {

	fmt.Println()
	fmt.Println("==================================================")
	fmt.Println("                   BLOCKCHAIN")
	fmt.Println("==================================================")

	for _, block := range blockchain {

		fmt.Println()
		fmt.Println("--------------------------------------------------")

		fmt.Println("BLOCO:", block.Index)

		fmt.Println("Data:", block.Timestamp)

		fmt.Println("Minerador:", block.Miner)

		fmt.Println()

		fmt.Println("TRANSAÇÕES:")

		// Mostramos todas as transações do bloco.
		for i, transaction := range block.Transactions {

			fmt.Println()
			fmt.Println("  Transação:", i+1)

			showTransaction(transaction)
		}

		fmt.Println()

		fmt.Println("Previous Hash:")
		fmt.Println(block.PreviousHash)

		fmt.Println()

		fmt.Println("Hash:")
		fmt.Println(block.Hash)

		fmt.Println()

		fmt.Println("Nonce:", block.Nonce)

		fmt.Println("Tentativas:", block.Attempts)

		fmt.Println(
			"Tempo de mineração:",
			block.MiningTime,
		)
	}

	fmt.Println()
	fmt.Println("==================================================")
}

// ============================================================
// MOSTRAR ÚLTIMO BLOCO
// ============================================================

func showLastBlock(blockchain []Block) {

	if len(blockchain) == 0 {

		fmt.Println("Blockchain vazia.")

		return
	}

	// len(blockchain)-1 representa a última posição.
	lastBlock := blockchain[len(blockchain)-1]

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("             ÚLTIMO BLOCO")
	fmt.Println("========================================")

	fmt.Println("Número:", lastBlock.Index)

	fmt.Println("Data:", lastBlock.Timestamp)

	fmt.Println("Minerador:", lastBlock.Miner)

	fmt.Println("Transações:", len(lastBlock.Transactions))

	fmt.Println("Nonce:", lastBlock.Nonce)

	fmt.Println("Tentativas:", lastBlock.Attempts)

	fmt.Println(
		"Tempo de mineração:",
		lastBlock.MiningTime,
	)

	fmt.Println()
	fmt.Println("Hash:")
	fmt.Println(lastBlock.Hash)

	fmt.Println("========================================")
}

// ============================================================
// MOSTRAR TRANSAÇÕES PENDENTES
// ============================================================

func showPendingTransactions(
	pendingTransactions []Transaction,
) {

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("        TRANSAÇÕES PENDENTES")
	fmt.Println("========================================")

	if len(pendingTransactions) == 0 {

		fmt.Println("Nenhuma transação aguardando mineração.")

		return
	}

	for i, transaction := range pendingTransactions {

		fmt.Println()
		fmt.Println("Transação:", i+1)

		showTransaction(transaction)
	}

	fmt.Println()
	fmt.Println(
		"Total:",
		len(pendingTransactions),
		"transação(ões)",
	)
}

// ============================================================
// INFORMAÇÕES DA BLOCKCHAIN
// ============================================================

func showBlockchainInfo(
	blockchain []Block,
	pendingTransactions []Transaction,
) {

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("      INFORMAÇÕES DA BLOCKCHAIN")
	fmt.Println("========================================")

	// Quantidade de blocos.
	fmt.Println("Quantidade de blocos:", len(blockchain))

	// Quantidade de transações confirmadas.
	totalTransactions := 0

	for _, block := range blockchain {

		totalTransactions += len(block.Transactions)
	}

	fmt.Println(
		"Transações armazenadas:",
		totalTransactions,
	)

	fmt.Println(
		"Transações pendentes:",
		len(pendingTransactions),
	)

	fmt.Println(
		"Dificuldade:",
		difficulty,
	)

	fmt.Println(
		"Arquivo do ledger:",
		ledgerFile,
	)

	// Mostramos o último bloco.
	if len(blockchain) > 0 {

		lastBlock := blockchain[len(blockchain)-1]

		fmt.Println(
			"Último bloco:",
			lastBlock.Index,
		)

		fmt.Println(
			"Último minerador:",
			lastBlock.Miner,
		)
	}

	fmt.Println("========================================")
}

// ============================================================
// LER VALOR FLOAT
// ============================================================

// readFloat é uma função auxiliar para ler números decimais.
func readFloat(scanner *bufio.Scanner) (float64, error) {

	scanner.Scan()

	text := strings.TrimSpace(scanner.Text())

	// Permite o aluno digitar:
	//
	// 150,50
	//
	// ou
	//
	// 150.50
	//
	text = strings.ReplaceAll(text, ",", ".")

	return strconv.ParseFloat(text, 64)
}

// ============================================================
// MAIN
// ============================================================

func main() {

	// Blockchain carregada na memória.
	var blockchain []Block

	// Transações que ainda não foram mineradas.
	var pendingTransactions []Transaction

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("        BLOCKCHAIN EDUCACIONAL GO")
	fmt.Println("========================================")

	// ========================================================
	// TENTAMOS CARREGAR O LEDGER
	// ========================================================

	loadedBlockchain, err := loadBlockchain()

	if err != nil {

		// ----------------------------------------------------
		// ARQUIVO NÃO EXISTE
		// ----------------------------------------------------

		if os.IsNotExist(err) {

			fmt.Println()
			fmt.Println("Ledger não encontrado.")
			fmt.Println("Esta parece ser a primeira execução.")
			fmt.Println()

			fmt.Println("Criando uma nova blockchain...")

			// Criamos o Genesis Block.
			genesisBlock := createGenesisBlock()

			// Adicionamos à blockchain.
			blockchain = append(
				blockchain,
				genesisBlock,
			)

			// Salvamos imediatamente.
			err = saveBlockchain(blockchain)

			if err != nil {

				fmt.Println(
					"Erro ao salvar blockchain:",
					err,
				)

				return
			}

			fmt.Println()
			fmt.Println(
				"Blockchain criada com sucesso!",
			)

		} else {

			// ------------------------------------------------
			// OUTRO ERRO
			// ------------------------------------------------

			fmt.Println(
				"Erro ao carregar o ledger:",
				err,
			)

			return
		}

	} else {

		// ====================================================
		// LEDGER ENCONTRADO
		// ====================================================

		blockchain = loadedBlockchain

		fmt.Println()
		fmt.Println("Ledger encontrado.")

		fmt.Println(
			"Blocos carregados:",
			len(blockchain),
		)

		// ====================================================
		// VALIDAÇÃO AUTOMÁTICA
		// ====================================================

		fmt.Println()
		fmt.Println(
			"Verificando integridade da blockchain...",
		)

		if !validateBlockchain(blockchain) {

			fmt.Println()
			fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			fmt.Println("ERRO CRÍTICO!")
			fmt.Println("A BLOCKCHAIN ESTÁ INVÁLIDA.")
			fmt.Println("O LEDGER PODE TER SIDO MODIFICADO.")
			fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			fmt.Println()
			fmt.Println(
				"O programa será encerrado para evitar",
			)
			fmt.Println(
				"que novos blocos sejam adicionados.",
			)

			return
		}

		fmt.Println()
		fmt.Println(
			"Blockchain válida!",
		)

		fmt.Println(
			"Continuaremos a partir do bloco:",
			blockchain[len(blockchain)-1].Index,
		)
	}

	// ========================================================
	// SCANNER PARA LER O TECLADO
	// ========================================================

	scanner := bufio.NewScanner(os.Stdin)

	// ========================================================
	// MENU PRINCIPAL
	// ========================================================

	for {

		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("          BLOCKCHAIN EM GO")
		fmt.Println("========================================")
		fmt.Println("1 - Criar nova transação")
		fmt.Println("2 - Minerar bloco")
		fmt.Println("3 - Mostrar blockchain")
		fmt.Println("4 - Validar blockchain")
		fmt.Println("5 - Mostrar último bloco")
		fmt.Println("6 - Informações da blockchain")
		fmt.Println("7 - Mostrar transações pendentes")
		fmt.Println("0 - Sair")
		fmt.Println("========================================")

		fmt.Print("Escolha uma opção: ")

		scanner.Scan()

		option := strings.TrimSpace(scanner.Text())

		// ====================================================
		// SWITCH DO MENU
		// ====================================================

		switch option {

		// ====================================================
		// OPÇÃO 1
		// CRIAR NOVA TRANSAÇÃO
		// ====================================================

		case "1":

			fmt.Println()
			fmt.Println("========================================")
			fmt.Println("           NOVA TRANSAÇÃO")
			fmt.Println("========================================")

			// Perguntamos quem envia.
			fmt.Print("Origem: ")

			scanner.Scan()

			from := strings.TrimSpace(scanner.Text())

			if from == "" {

				fmt.Println(
					"Origem não pode ficar vazia.",
				)

				continue
			}

			// Perguntamos quem recebe.
			fmt.Print("Destino: ")

			scanner.Scan()

			to := strings.TrimSpace(scanner.Text())

			if to == "" {

				fmt.Println(
					"Destino não pode ficar vazio.",
				)

				continue
			}

			// Perguntamos o valor.
			fmt.Print("Valor: R$ ")

			amount, err := readFloat(scanner)

			if err != nil {

				fmt.Println(
					"Valor inválido.",
				)

				continue
			}

			// Não permitimos valor zero ou negativo.
			if amount <= 0 {

				fmt.Println(
					"O valor precisa ser maior que zero.",
				)

				continue
			}

			// Criamos a transação.
			transaction := Transaction{
				From:   from,
				To:     to,
				Amount: amount,
			}

			// IMPORTANTE:
			//
			// A transação ainda NÃO está na blockchain.
			//
			// Ela fica em uma lista de transações pendentes
			// esperando que alguém mine um bloco.
			pendingTransactions = append(
				pendingTransactions,
				transaction,
			)

			fmt.Println()
			fmt.Println(
				"Transação criada com sucesso!",
			)

			fmt.Println()
			fmt.Println(
				"Ela está aguardando mineração.",
			)

			fmt.Println(
				"Transações pendentes:",
				len(pendingTransactions),
			)

		// ====================================================
		// OPÇÃO 2
		// MINERAR BLOCO
		// ====================================================

		case "2":

			// Não faz sentido minerar se não existem
			// transações.
			if len(pendingTransactions) == 0 {

				fmt.Println()
				fmt.Println(
					"Não existem transações pendentes.",
				)

				fmt.Println(
					"Crie uma transação antes de minerar.",
				)

				continue
			}

			fmt.Println()
			fmt.Println("========================================")
			fmt.Println("             MINERAÇÃO")
			fmt.Println("========================================")

			// Perguntamos quem está minerando.
			fmt.Print("Nome do minerador: ")

			scanner.Scan()

			miner := strings.TrimSpace(scanner.Text())

			if miner == "" {

				fmt.Println(
					"O nome do minerador não pode ficar vazio.",
				)

				continue
			}

			// Pegamos o último bloco da blockchain.
			lastBlock := blockchain[len(blockchain)-1]

			// Criamos uma CÓPIA das transações pendentes.
			//
			// Isso é importante porque elas serão incluídas
			// no novo bloco.
			transactionsToMine := make(
				[]Transaction,
				len(pendingTransactions),
			)

			copy(
				transactionsToMine,
				pendingTransactions,
			)

			// Criamos e mineramos o novo bloco.
			newBlock := createBlock(
				transactionsToMine,
				miner,
				lastBlock,
			)

			// Adicionamos o bloco à blockchain.
			blockchain = append(
				blockchain,
				newBlock,
			)

			// Agora salvamos o ledger no disco.
			err := saveBlockchain(blockchain)

			if err != nil {

				fmt.Println()
				fmt.Println(
					"Erro ao salvar blockchain:",
					err,
				)

				// Como houve erro no salvamento,
				// não apagamos as transações pendentes.
				continue
			}

			// O bloco foi salvo corretamente.
			//
			// Portanto podemos limpar as transações
			// pendentes.
			pendingTransactions = nil

			fmt.Println()
			fmt.Println("========================================")
			fmt.Println("BLOCO ADICIONADO À BLOCKCHAIN!")
			fmt.Println("========================================")

			fmt.Println(
				"Número do bloco:",
				newBlock.Index,
			)

			fmt.Println(
				"Transações confirmadas:",
				len(newBlock.Transactions),
			)

			fmt.Println(
				"Minerador:",
				newBlock.Miner,
			)

		// ====================================================
		// OPÇÃO 3
		// MOSTRAR BLOCKCHAIN
		// ====================================================

		case "3":

			showBlockchain(blockchain)

		// ====================================================
		// OPÇÃO 4
		// VALIDAR BLOCKCHAIN
		// ====================================================

		case "4":

			fmt.Println()
			fmt.Println(
				"Verificando blockchain...",
			)

			if validateBlockchain(blockchain) {

				fmt.Println()
				fmt.Println(
					"BLOCKCHAIN VÁLIDA!",
				)

				fmt.Println(
					"Nenhuma adulteração foi detectada.",
				)

			} else {

				fmt.Println()
				fmt.Println(
					"BLOCKCHAIN INVÁLIDA!",
				)

				fmt.Println(
					"Algum bloco foi modificado",
				)

				fmt.Println(
					"ou a cadeia foi quebrada.",
				)
			}

		// ====================================================
		// OPÇÃO 5
		// MOSTRAR ÚLTIMO BLOCO
		// ====================================================

		case "5":

			showLastBlock(blockchain)

		// ====================================================
		// OPÇÃO 6
		// INFORMAÇÕES
		// ====================================================

		case "6":

			showBlockchainInfo(
				blockchain,
				pendingTransactions,
			)

		// ====================================================
		// OPÇÃO 7
		// TRANSAÇÕES PENDENTES
		// ====================================================

		case "7":

			showPendingTransactions(
				pendingTransactions,
			)

		// ====================================================
		// OPÇÃO 0
		// SAIR
		// ====================================================

		case "0":

			fmt.Println()
			fmt.Println(
				"Encerrando a blockchain...",
			)

			// Avisamos caso existam transações que
			// ainda não foram mineradas.
			if len(pendingTransactions) > 0 {

				fmt.Println()
				fmt.Println(
					"ATENÇÃO:",
				)

				fmt.Println(
					len(pendingTransactions),
					"transação(ões) não foram mineradas.",
				)

				fmt.Println(
					"Elas serão perdidas ao fechar o programa.",
				)
			}

			fmt.Println()
			fmt.Println(
				"Ledger salvo em:",
				ledgerFile,
			)

			fmt.Println(
				"Programa encerrado.",
			)

			return

		// ====================================================
		// OPÇÃO INVÁLIDA
		// ====================================================

		default:

			fmt.Println()
			fmt.Println(
				"Opção inválida.",
			)

			fmt.Println(
				"Escolha uma opção entre 0 e 7.",
			)
		}
	}
}
