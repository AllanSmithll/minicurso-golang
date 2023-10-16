/*
Allan Alves Amâncio - Minicurso de Golang - IFTECH - IFPB
Conteúdo: Manipulação de Strings
Data: 16/10/2023
*/
package main

import ("fmt", 
			"strings" )

func main() {
	// Declaração de strings
	texto := "Isso é uma string"

	// Concatenação de strings
	string1 := "Olá"
	string2 := " Mundo"
	resultado := string1 + string2 // Concatenação com operador '+'

	// Comprimento de strings
	texto := "Golang"
	tamanho := len(texto) // Tamanho da string

	// Índices e Acesso a Caracteres
	texto1 := "Golang"
	primeiroCaracter := texto1[0] // 'G'

	// Iteração em Strings
	texto2 := "Golang"
	for i := 0; i < len(texto2); i++ {
    	fmt.Printf("%c ", texto2[i])
	}

	// Funções do Pacote “strings”
	texto := "Esta é uma frase de exemplo"
	palavras := strings.Split(texto, " ") // Divide a string em palavras

	// Comparação de Strings
	str1 := "abc"
	str2 := "abc"
	igual := str1 == str2 // true

	// Formatação de Strings
	nome := "Alice"
	idade := 30
	frase := fmt.Sprintf("Nome: %s, Idade: %d", nome, idade)
}