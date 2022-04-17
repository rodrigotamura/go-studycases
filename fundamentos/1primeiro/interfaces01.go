package main

import "fmt"

type imprimivel interface {
	toString() string
	// exemplo: metodo q toda interface imprimivel deve ter
	// para saber que uma estrutura é imprimivel, basta ter este método
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
// em Go não veremos que pessoa implementa imprimivel
// basta criarmos uma struct que RESPEITE a interface criada

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

// aqui podemos definir que o x DEVE seguir a interface tipada
func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	// podemos tipar a variável com a interface, porém o struct
	// deverá seguir a interface criada (precisa ter o toString())
	var coisa imprimivel = pessoa{nome: "Rodrigo", sobrenome: "Tamura"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	// polimorfismo - a coisa (pessoa) virando outra coisa (produto)
	coisa = produto{nome: "Calça", preco: 19.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	// respeitando a interface, podemos passar diretamente o produto
	imprimir(produto{nome: "Calça", preco: 19.90})
}
