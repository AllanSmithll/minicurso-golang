/*
Allan Alves Amâncio - Minicurso de Golang - IFTECH - IFPB
Conteúdo: Exercício 6 sobre slices
Data: 16/10/2023
*/
package main

import "fmt"

func main() {
	var frutas_iniciais [5]string;
	frutas_iniciais = [5]string{"maça", "manga", "uva", "morango", "melancia"};  // Array do exercício 5

	mais_frutas := append(frutas_iniciais[:]); // Slice de frutas
	mais_frutas = append(mais_frutas, "pera");
	mais_frutas = append(mais_frutas, "goiaba");
	mais_frutas = append(mais_frutas, "laranja");

	fmt.Println(mais_frutas);
}