package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

// Método para envelhecer uma pessoa
func (p *Pessoa) Envelhecer() {
	p.Idade++
}

// Função para imprimir informações da pessoa
func ImprimirPessoa(p Pessoa) {
	fmt.Printf("Nome: %s, Idade: %d\n", p.Nome, p.Idade)
}

func main() {
	var pessoa Pessoa
	var nome string;
	pessoa = Pessoa{"Alice", 25}

	nome = pessoa.Nome // Acessar o campo “Nome“
	pessoa.Idade = 30   // Modificar o campo “Idade“

	alice := Pessoa{"Alice", 25}
	ImprimirPessoa(alice) // Nome: Alice, Idade: 25

	fmt.Println(nome)
}