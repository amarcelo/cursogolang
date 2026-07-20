package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	// Frase original
	frase := "aprendendo gerenciamento de dependências em go"

	// Cria um conversor para letras maiúsculas
	conversor := cases.Upper(language.Portuguese)

	// Converte a frase
	fraseMaiuscula := conversor.String(frase)

	// Mostra o resultado
	fmt.Println("Frase original:")
	fmt.Println(frase)

	fmt.Println("\nFrase em letras maiúsculas:")
	fmt.Println(fraseMaiuscula)
}
