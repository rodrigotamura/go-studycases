package main

import (
	"fmt"
	"time"
)

type item struct {
	nome       string
	valor      float64
	quantidade int64
}

type pedido struct {
	numero int64
	data   string
	itens  []item
}

func (p pedido) calculaTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.valor * float64(item.quantidade)
	}

	p.numero = 444

	return total
}

func main() {
	pedido1 := pedido{
		numero: 1,
		data:   time.Now().Local().String(),
		itens: []item{
			item{nome: "xampu", valor: 5.99, quantidade: 3},
			item{nome: "sabonete", valor: 0.99, quantidade: 2},
		},
	}

	fmt.Println(pedido1)
	fmt.Println(pedido1.calculaTotal())
	fmt.Println(pedido1)
}
