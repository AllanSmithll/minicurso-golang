/*
Allan Alves Amâncio - Minicurso de Golang - IFTECH - IFPB
Conteúdo: Variáveis e Constantes
Data: 12/10/2023
*/
package main

import "fmt"

// %d -> inteiro
// %f -> float
// %s -> string
// %c -> caracter

func main() {
	var idade = 18                                               // Número inteiro variável
	const PI = 3.14159                                           // Número "flutuante" constante
	var resultado = fmt.Sprintf("Idade: %d - Pi: %f", idade, PI) // Concatenando strings
	fmt.Println(resultado)                                       // Mostrando as variáveis
}
