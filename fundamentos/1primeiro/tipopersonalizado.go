package main

import "fmt"

type nota float64

func (n nota) mediaNota(mediaNota float64) {
	if float64(n) > float64(mediaNota) {
		fmt.Println("Acima da média")
	} else {
		fmt.Println("Abaixo da média")
	}
}

func passei(pontos nota) {
	fmt.Printf("Nota atual: %.2f", pontos)
	pontos.mediaNota(6.0)
}

func main() {
	passei(5)
}
