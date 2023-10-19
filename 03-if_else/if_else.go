/*
Allan Alves Amâncio - Minicurso de Golang - IFTECH - IFPB
Conteúdo: Estruturas de Controle
Data: 12/10/2023
*/
package main

import "fmt"

func main() {
	const IDADE = 16
	if IDADE >= 18 {
		fmt.Println("Você é de maior!")
	} else {
		fmt.Println("Você não é de maior!")
	}
}
