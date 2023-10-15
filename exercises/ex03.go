/*
Allan Alves Amâncio - Minicurso de Golang - IFTECH - IFPB
Conteúdo: Exercício 3 sobre estrutura de repetição
Data: 15/10/2023
*/
package main

import "fmt"

func main() {
    // Solicite ao usuário que insira um número inteiro N
    fmt.Print("Digite um número inteiro N: ")
    var N int
    fmt.Scan(&N)

    // Use um loop for para imprimir os números pares de 1 a N
    fmt.Printf("Números pares de 1 a %d:\n", N)
    for i := 1; i <= N; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }
}