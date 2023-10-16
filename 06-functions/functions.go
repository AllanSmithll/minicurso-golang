/*
Allan Alves Amâncio - Minicurso de Golang - IFTECH - IFPB
Conteúdo: Funções
Data: 15/10/2023
*/
package main

import "fmt"

func soma(a int, b int) int {
	resultado := a + b
	return resultado
}

func main() {
	var result = soma(2, 4)
	fmt.Println("A soma dos números é igual a", result)
}
