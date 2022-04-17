package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Realizando requisicao para server 2...")

	req, err := http.NewRequest("GET", "http://localhost:8001", nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))

	if err != nil {
		fmt.Errorf("erro: %v", err)
	}
	// Be sure to handle errors.
	defer resp.Body.Close()

	fmt.Printf("Resposta do server 2: %s \n\n", resp.Body)

	w.Write([]byte(fmt.Sprintf("Response from server 2: %s", resp.Body)))
}
