/*
Allan Alves Amâncio - Minicurso de Golang - IFTECH - IFPB
Conteúdo: Exercício 1 sobre introdução ao Golang
Data: 15/10/2023
*/
package main

import "fmt"

func main() {
	const nome string = "Allan";
	const idade int = 18;
	const curso = "TSI";
	const minicurso = "Conceitos Fundamentais do Golang";
	fmt.Printf(`
	Nome: %s
	Idade: %d
	Curso: %s
	Minicurso: %s
	`, nome, idade, curso, minicurso)
}