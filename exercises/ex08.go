/*
Allan Alves Amâncio - Minicurso de Golang - IFTECH - IFPB
Conteúdo: Exercício 8 sobre manipulação de strings
Data: 16/10/2023
*/
package main

import "fmt"

func inverterFrase(frase string) string {
    resultado := ""
    for i := len(frase) - 1; i >= 0; i-- {
        resultado += string(frase[i])
    }
    return resultado
}

func main() {
    frase := "O Go e incrivelmente ...!"
    fraseInvertida := inverterFrase(frase)

    fmt.Printf("Frase original: %s\n", frase)
    fmt.Printf("Frase invertida: %s\n", fraseInvertida)
}