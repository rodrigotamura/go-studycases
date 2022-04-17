package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init do main")
}

func funcaoMain() {
	fmt.Println("Funcao qualquer do main")
}

func main() {
	fmt.Println("Executando main")

	// funcao()
	funcaoMain()
}
