package main

/*
Criando um HTTP server estático,
servindo páginas estáticas
dinamicamente
*/

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public")) // fileserver para ler o diretorio

	http.Handle("/", fs)

	log.Println("Ëxecutando...:")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
