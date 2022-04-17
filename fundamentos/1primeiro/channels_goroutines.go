package main

import (
	"fmt"
	"time"
)

func doiTresQuatroVezes(n int, ch chan int) {
	time.Sleep(time.Second)
	ch <- 2 * n // enquanto esse canal não for lido - <-c - não prossegue adiante

	time.Sleep(time.Second)
	ch <- 3 * n

	time.Sleep(3 * time.Second)
	ch <- 4 * n
}

func main() {
	c := make(chan int)

	go doiTresQuatroVezes(2, c)

	fmt.Println("Aqui imprime imediatamente msm dps do goroutine")

	// então aqui estamos declarando 2x o <-c, significando que dois pontos de canais
	// serão permitidos prosseguirem
	a, b := <-c, <-c

	fmt.Println(a, b)
	fmt.Println("Aqui imprime só quando o primeiro e o segundo channel resolver")

	fmt.Println(<-c)
	fmt.Println("Finaliza somente que o terceiro channel resolver")
}
