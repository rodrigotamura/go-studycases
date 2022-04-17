package main

import "fmt"

func main() {
	ch := make(
		chan int64, // criando uma channel que vai trafegar dados int64
		1,          //buffer - veremos dps
	)

	ch <- 1 // enviando o valor 1 para o canal ch (escrita)
	<-ch    // recebendo dados do canal ch

	ch <- 2
	fmt.Println(<-ch)

}
