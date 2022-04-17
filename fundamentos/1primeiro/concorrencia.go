package main

import (
	"fmt"
	"time"
)

func fale(nome string, texto string, quantidade int) {
	for i := 0; i < quantidade; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s fala: \"%s\" (iteracao: %d) \n", nome, texto, i)
	}
}

func main() {
	go fale("Rodrigo", "Olá", 3)
	go fale("Debora", "Olá depois do Rodrigo", 1)
}
