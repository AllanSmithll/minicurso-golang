/*
Allan Alves Amâncio - Minicurso de Golang - IFTECH - IFPB
Conteúdo: Exercício 7 sobre structs
Data: 16/10/2023
*/
package main

import "fmt"

type Animal struct {
    Nome           string
    Idade          int
    Raca           string
    Peso           float64
    Data_Nascimento string
    Origem         string
}

func main() {
    zoo := []Animal{} // Slice da struct Animal

    animal1 := Animal{"Leão", 5, "Africano", 190.5, "10/05/2018", "Savana"}
    animal2 := Animal{"Elefante", 10, "Africano", 4000, "02/03/2010", "Savana"}
    zoo = append(zoo, animal1)
    zoo = append(zoo, animal2)

    fmt.Println("Animais no Zoológico:")
    for _, animal := range zoo {
        fmt.Printf("Nome: %s, Idade: %d, Raça: %s, Peso: %.2f kg, Data de Nascimento: %s, Origem: %s\n", animal.Nome, animal.Idade, animal.Raca, animal.Peso, animal.Data_Nascimento, animal.Origem)
    }

    nomeAnimalARemover := "Leão"
    for i, animal := range zoo {
        if animal.Nome == nomeAnimalARemover {
            zoo = append(zoo[:i], zoo[i+1:]...)
            break
        }
    }

    fmt.Println("\nApós Remoção:")
    for _, animal := range zoo {
        fmt.Printf("Nome: %s, Idade: %d, Raça: %s, Peso: %.2f kg, Data de Nascimento: %s, Origem: %s\n", animal.Nome, animal.Idade, animal.Raca, animal.Peso, animal.Data_Nascimento, animal.Origem)
    }
}