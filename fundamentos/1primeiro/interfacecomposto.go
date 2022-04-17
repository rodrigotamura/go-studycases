package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// pode adicionar novos métodos
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo ligado")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Fazer Baliza!")
}

func main() {
	var carro esportivoLuxuoso = bmw7{}
	carro.fazerBaliza()
	carro.ligarTurbo()
}
