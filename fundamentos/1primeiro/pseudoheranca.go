package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campo anonimo, onde faremos a pseudo herança
	turboLigado bool
}

func (f *ferrari) mudaNome(nome string) {
	f.nome = nome
}

func main() {
	f := ferrari{}

	f.nome = "F90"

	fmt.Print(f)

	f.mudaNome("F40")

	fmt.Print(f)
}
