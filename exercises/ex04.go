/*
Allan Alves Amâncio - Minicurso de Golang - IFTECH - IFPB
Conteúdo: Exercício 4 sobre funções
Data: 15/10/2023
*/
package main

import "fmt"

// Função para verificar se um número é par ou ímpar
func verificarParOuImpar(numero int) string {
	if numero%2 == 0 {
		return "par"
	} else {
		return "ímpar"
	}
}

func main() {
	// Solicite ao usuário que insira um número inteiro
	fmt.Print("Digite um número inteiro: ")
	var numero int
	fmt.Scan(&numero)

	// Chame a função para verificar se o número é par ou ímpar e imprima o resultado
	resultado := verificarParOuImpar(numero)
	fmt.Printf("O número %d é %s.\n", numero, resultado)
}
