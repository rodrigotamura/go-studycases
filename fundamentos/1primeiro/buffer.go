package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Executou até aqui respeitando o limite de buffer=3")
	ch <- 4
	fmt.Println("Agora que esvaziou o buffer, veio aqui tbm")
	ch <- 5
	ch <- 6
	ch <- 7
}

func main() {
	ch := make(chan int, 3) // buffer de 3 posicoes

	go rotina(ch)

	time.Sleep(time.Second * 2)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("Buffer tá vazio!")
	fmt.Println(<-ch)
}
