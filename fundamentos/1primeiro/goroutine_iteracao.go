package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

// n: quantidade de numeros primos a retornar
// c: channel
func primos(n int, c chan int) {
	inicio := 2

	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Second)
				break
			}
		}
	}

	/*
		c <- 1
		time.Sleep(time.Second)
		c <- 2
		time.Sleep(time.Second)
		c <- 3
		time.Sleep(time.Second)
		c <- 4
		time.Sleep(time.Second)
		c <- 5
	*/
	close(c) // dentro desse canal de comunicacao nao vai ser enviado mais nenhum dado
	// se não colocarmos close() vai ser gerado um deadlock
}

func main() {
	c := make(chan int, 30)
	go primos(
		cap(c), // funcao q retorna o tamanho do buffer do canal - 30
		c,      // canal em si
	)

	for primo := range c { // iteracao de channel
		fmt.Printf("%d ", primo)
	}

	fmt.Println("Fim!")
}
