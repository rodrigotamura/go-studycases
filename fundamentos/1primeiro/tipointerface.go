package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{} // declarando do tipo interface vazia
	fmt.Println(coisa)    // nil

	coisa = 3
	fmt.Println(coisa) // 3

	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2) // Opa

	coisa2 = 1
	fmt.Println(coisa2) // 1

	coisa2 = curso{"Golang"}
	fmt.Println(coisa2) // {Golang}
}
