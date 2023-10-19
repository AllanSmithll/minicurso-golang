/*
Allan Alves Amâncio - Minicurso de Golang - IFTECH - IFPB
Conteúdo: Exercício 5 sobre arrays
Data: 16/10/2023
*/
package main

import "fmt"

func main() {
	var frutas [5]string
	frutas = [5]string{"maça", "manga", "uva", "morango", "melancia"}

	fruta_escolhida := "maça"

	for i, fruta := range frutas {
		if fruta == fruta_escolhida {
			fmt.Printf("A fruta escolhida (%s) está no array, na posição %d.", frutas[i], i)
		}
	}
}
