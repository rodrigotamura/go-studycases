package main

import (
	"fmt"
	"time"
)

// vamos utilizar aqui as funcoes de obtencao de titulos por url
// ao qual aplicamos o padrão generators para encapsular o goroutine

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	// estrutura de controle para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		// caso haja um timeout de 1 segundo
		return "Todos perderam"
		// se executarmos com o default descomentado, como existe um delay para obtencao
		// de alguma resposta acima, imediatamente cairá aqui. Por isso vamos comentar.
		// default:
		// 	return "Sem resposta ainda"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://google.com.br",
		"https://www.amazon.com",
		"https://www.uol.com.br",
	)

	fmt.Println(campeao)
}
