package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

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

func main() {
	// vamos criar dois channels
	t1 := titulo(
		"https://www.youtube.com.br",
		"https://www.ifood.com.br",
	)

	t2 := titulo(
		"https://www.g1.com.br",
		"https://www.uol.com.br",
	)

	fmt.Println("Primeiros: ", <-t1, " | ", <-t2)
	fmt.Println("Segundos: ", <-t1, " | ", <-t2)

}
