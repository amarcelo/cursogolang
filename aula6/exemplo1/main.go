package main

// Importa o pacote fmt para imprimir informações na tela
import "fmt"

func main() {

	// Cria um slice de inteiros
	//
	// []int  -> slice que armazenará números inteiros
	// 0      -> tamanho inicial (não possui elementos)
	// 5      -> capacidade inicial (espaço reservado para 5 elementos)
	//
	// Visualmente:
	// [ _ _ _ _ _ ]
	//
	// len = 0
	// cap = 5
	numeros := make([]int, 0, 5)

	// Mostra o conteúdo do slice
	fmt.Println("Slice inicial:", numeros)

	// len() mostra quantos elementos existem no slice
	fmt.Println("Tamanho:", len(numeros))

	// cap() mostra a capacidade total reservada
	fmt.Println("Capacidade:", cap(numeros))

	// Adiciona o número 10 ao slice
	//
	// Antes:
	// [ _ _ _ _ _ ]
	//
	// Depois:
	// [10 _ _ _ _ ]
	numeros = append(numeros, 10)

	// Adiciona o número 20
	//
	// [10 20 _ _ _ ]
	numeros = append(numeros, 20)

	// Adiciona o número 30
	//
	// [10 20 30 _ _ ]
	numeros = append(numeros, 30)
	numeros = append(numeros, 40)
	numeros = append(numeros, 50)
	numeros = append(numeros, 60)

	fmt.Println("\nApós append:")

	// Exibe todos os elementos do slice
	fmt.Println("Slice:", numeros)

	// Agora existem 3 elementos
	fmt.Println("Tamanho:", len(numeros))

	// Ainda há espaço para até 5 elementos
	fmt.Println("Capacidade:", cap(numeros))

	fmt.Println("\nElementos do slice:")

	// range percorre todos os elementos do slice
	//
	// i     -> posição do elemento
	// valor -> valor armazenado naquela posição
	for i, valor := range numeros {

		// Imprime a posição e o valor
		fmt.Printf("Posição %d = %d\n", i, valor)
	}
}
