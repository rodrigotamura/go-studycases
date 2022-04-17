package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8001", nil)
}

func otherHandler(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(4 * time.Second):
			// time's up
		case <-r.Context().Done():
			// client gave up
			return
		}

		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
		fmt.Println("response:", "The time is: "+tm)
	}
}

func Pausada(c chan string) {
	time.Sleep(time.Second * 10)
	c <- "foi"
}

func Carregando(canal chan string) string {
	fmt.Println("Carregando:")
	for {
		time.Sleep(time.Second)
		select {
		case t1 := <-canal:
			fmt.Printf("\n")
			return t1
		default:
			fmt.Printf("|")
		}
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string)

	fmt.Println("Iniciando server 2") // output aplicacao

	go Pausada(ch)

	Carregando(ch)

	fmt.Fprintf(w, "Passaram 10 segundos!") // output response API

	fmt.Println("\nServer 2 finalizado!")

	select {
	case <-r.Context().Done():
		fmt.Errorf("cliente desistiu")
		return
	}
}
