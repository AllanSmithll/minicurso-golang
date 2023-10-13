/*
Allan Alves Amâncio - Minicurso de Golang - IFTECH - IFPB
Conteúdo: Loops (for)
Data: 12/10/2023
*/
package main

import "fmt"

func main() {
	const MAX_ROTACOES_LACO = 10 // Máximo de rotações do laço
	for i := 1; i <= MAX_ROTACOES_LACO; i++ {
		fmt.Println("Número:", i) // é mostrado na tela os números de 1 até o máximo de rotações do laço
	}
}
