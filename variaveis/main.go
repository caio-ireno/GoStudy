package main

import (
	"fmt"
	"gostudy/variaveis/matematica"
)



func main() {
	// Declaração explicita
	var nome string = "Caio"

	// Declaração de inferencia de tipo
	var idade = 25

	// Comum e usado quando você quer algo rápido e o tipo pode ser inferido.
	cidade := "São Paulo"

	altura := 1.80

	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Cidade:", cidade)
	fmt.Println("Altura:", altura)

	fmt.Println("------ VALUE ---------")
	fmt.Printf("Nome: %v \n", nome)
	fmt.Printf("Idade: %v \n", idade)
	fmt.Printf("Cidade: %v \n", cidade)
	fmt.Printf("Altura: %v \n", altura)

	fmt.Println("------ TYPE ---------")
	fmt.Printf("Nome: %T \n", nome)
	fmt.Printf("Idade: %T \n", idade)
	fmt.Printf("Cidade: %T \n", cidade)
	fmt.Printf("Altura: %T \n", altura)

	fmt.Println("------ FUNÇÃO SOMA ---------")
	a:= 32
	b:= 10

	var soma int = matematica.Soma(a, b)
	//soma := Soma(a, b)
	fmt.Println("Soma:", soma)
	fmt.Printf("Soma: %T \n", soma)

	fmt.Println("------ FUNÇÃO SOMA FIXA ---------")
	somaFixa := matematica.SomaFixa20(a)
	fmt.Println("Soma Fixa:", somaFixa)

	var internal = matematica.InternalVar
	fmt.Println("InternalVar:", internal)
}

