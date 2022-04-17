package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("01/01/2001 01:01:01")

	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s) // vai escrever dentro do http.ResponseWriter para ter output como response
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

type Usuario struct {
	id   int    `json="id"`
	nome string `json="nome"`
}

func anyHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3061)/bancodedados")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Erro inexperado")
		log.Fatalf("Erro de conexão: %s", err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios")
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.id, &usuario.nome)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}
