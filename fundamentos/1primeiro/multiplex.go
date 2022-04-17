package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// script da obtencao de titulo por url
// <-chan - canal somente leitura
func titulo(urls ...string) <-chan string { // retorna um canal somente leitura
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>") // procurando tag title

			// colocando no canal o valor capturado na regex criada acima
			// a partir do html obtido na url acessada via get
			c <- r.FindStringSubmatch(string(html))[1]
		}(url) // disparando imediatamente IIFL
	}
	return c // retorna mesmo que o fluxo ainda não concluiu, pois o que
	// queremos retornar é um canal
}

// fim do script da obtencao de titulo por html

func encaminhar(origem <-chan string, destino chan string) {
	// origem é do tipo channel de leitura (<-chan)
	// destino é do tipo leitura e escrita

	for {
		destino <- <-origem // sempre que chegar algo na origem, manda para o destino
		// enquanto tiver dados na origem continua encaminhando
		// quando não houver dados entra o fator bloqueante
	}
}

// multiplexar - misturar - mensagens - num único canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	// informa dois canais diferentes e retorna um canal de leitura
	c := make(chan string)

	go encaminhar(entrada1, c) // encaminhar entrada para canal c
	go encaminhar(entrada2, c) // encaminhar entrada para canal c

	return c
}

func main() {
	c := juntar(
		titulo(
			"https://www.youtube.com.br",
			"https://www.ifood.com.br",
		),
		titulo(
			"https://www.g1.com.br",
			"https://www.uol.com.br",
		),
	)

	fmt.Println(<-c) // o primeiro que chegar dos quatro, vai aparecer aqui
	fmt.Println(<-c) // o segundo que chegar dos quatro, vai aparecer aqui
	fmt.Println(<-c) // o terceiro que chegar dos quatro, vai aparecer aqui
	fmt.Println(<-c) // o quarto que chegar dos quatro, vai aparecer aqui
}
