/*
Allan Alves Amâncio - Minicurso de Golang - IFTECH - IFPB
Conteúdo: Exercício 2 sobre estruturas de decisão
Data: 15/10/2023
*/
package main

import "fmt"

func main() {
    // Solicite ao usuário que insira um número
    fmt.Print("Digite um número:")
    var numero int
    fmt.Scan(&numero)

    // Use uma estrutura de decisão para avaliar o número
    if numero > 0 {
        fmt.Println("O número é positivo.")
    } else if numero < 0 {
        fmt.Println("O número é negativo.")
    } else {
        fmt.Println("O número é igual a zero.")
    }
}