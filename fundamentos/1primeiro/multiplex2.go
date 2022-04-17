package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i <= 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d vezes", pessoa, i)
		}
	}()

	return c
}

// multiplexando dois canais em uma única saída
// através da "escuta" do canal 1 e 2
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for { // for inifinito
			select {
			case s := <-entrada1:
				c <- s // jogando para o c
			case s := <-entrada2:
				c <- s
			}
		}
	}()
	close(c)
	return c
}

func main() {
	c := juntar(falar("Rodrigo"), falar("Debora"))

	for fala := range c {
		fmt.Println(fala)
	}
}
