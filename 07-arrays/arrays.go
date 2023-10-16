/*
Allan Alves Amâncio - Minicurso de Golang - IFTECH - IFPB
Conteúdo: Arrays
Data: 16/10/2023
*/
package main

import "fmt"

func main() {
	var vogais [5]string;
	var vogais1 [5]string;
	var primeiroElemento string;
	var segundoElemento string;
	var tamanho int;
	var iguais bool;
	var slice []string;
	vogais = [5]string{"A", "E", "I", "O", "U"};
	vogais1 = [5]string {"A", "E", "I", "O", "ÃO"};
	fmt.Println(vogais);

	primeiroElemento = vogais[0]  // Acessa o primeiro elemento (índice 0)
	segundoElemento = vogais[1]  // Acessa o segundo elemento (índice 1)
	vogais[2] = "ÃO"  // Modifica o terceiro elemento (índice 2) para 42
	tamanho = len(vogais)  // Obtém o tamanho do array "vogais"
	iguais = vogais == vogais1  // Verifica se os arrays "vogais" e "consoantes" são iguais
	slice = vogais[1:4]  // Cria um slice que contém os elementos do índice 1 ao índice 3
}