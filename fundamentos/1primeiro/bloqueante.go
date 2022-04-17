package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)

	c <- 1 // ação bloqueante - enquanto ninguém lê, trava aqui
	fmt.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) // canal sem buffer - veremos dps

	go rotina(c)

	fmt.Println(<-c) // operação bloqueante. receberemos quando o dado estiver pronto
	fmt.Println("Foi lido")
	fmt.Println(<-c) // gerando deadlock
	fmt.Println("Fim! ")

}
